package product

import (
	"context"
	"database/sql"
	"fmt"
	interfaces "github.com/Torebekov/shop-back/internal/interfaces/repository"
	"github.com/Torebekov/shop-back/internal/models"
	"github.com/Torebekov/shop-back/modules/logger"
	"go.uber.org/zap"
	"log"
	"strings"
)

type product struct {
	db  *sql.DB
	ctx context.Context
}

func New(db *sql.DB, ctx context.Context) interfaces.IProduct {
	return &product{
		db:  db,
		ctx: ctx,
	}
}

func (r *product) List(searchText string, categoryID, userID uint64) (products []models.ProductDTO, err error) {
	l := logger.WorkLogger.Named("repo.product.List").With(zap.String("searchText", searchText),
		zap.Uint64("categoryID", categoryID))

	query := fmt.Sprintf(`
		SELECT
			product.id,
			product.name,
			category.name,
			product.image,
			product.price,
			(CASE WHEN user_favorite.product_id IS NULL THEN FALSE ELSE TRUE END) as is_favorite
		FROM
			product
		LEFT JOIN
			user_favorite ON user_favorite.product_id = product.ID AND user_id = %d
		JOIN
			category ON category.id = product.category_id`, userID)

	addWhereClause := func(condition string, value interface{}) {
		if value != nil {
			if strings.Contains(query, "WHERE") {
				query += " AND"
			} else {
				query += " WHERE"
			}

			query += fmt.Sprintf(condition, value)
		}
	}

	addWhereClause(" product.name LIKE '%%%s%%'", searchText)

	if categoryID != 0 {
		addWhereClause(" category_id = %d", categoryID)
	}

	log.Println(query)
	rows, err := r.db.Query(query)
	if err != nil {
		l.Error("couldn't make request", zap.Error(err))
		return
	}
	defer rows.Close()

	for rows.Next() {
		var productModel models.ProductDTO
		err = rows.Scan(&productModel.ID, &productModel.Name, &productModel.CategoryName, &productModel.Image, &productModel.Price, &productModel.IsFavorite)
		if err != nil {
			l.Error("couldn't scan product", zap.Error(err))
			return
		}
		log.Println(productModel)

		products = append(products, productModel)
	}

	if err = rows.Err(); err != nil {
		l.Error("iteration error", zap.Error(err))
		return
	}

	log.Println(products)

	return
}

func (r *product) ByID(productID uint64) (productModel models.Product, err error) {
	l := logger.WorkLogger.Named("repo.product.List")

	if r.db == nil {
		l.Error("DB not initialized")
		return
	}

	query := `
		SELECT
			id,
			name,
			category_id,
			image,
			price
		FROM
			product
			WHERE product.id = $1`

	row := r.db.QueryRow(query, productID)

	err = row.Scan(&productModel.ID, &productModel.Name, &productModel.CategoryID, &productModel.Image, &productModel.Price)
	if err != nil {
		l.Error("couldn't scan product", zap.Error(err))
		return
	}

	return
}

func (r *product) Add(productModel models.Product) (err error) {
	l := logger.WorkLogger.Named("repo.product.Add").With(zap.Any("productModel", productModel))

	if r.db == nil {
		l.Error("DB not initialized")
		return
	}

	_, err = r.db.Exec("INSERT INTO product (name, category_id, image, price) VALUES ($1, $2, $3, $4)", productModel.Name, productModel.CategoryID, productModel.Image, productModel.Price)
	if err != nil {
		l.Error("couldn't insert product")
		return
	}

	return
}

func (r *product) Remove(productID uint64) (err error) {
	l := logger.WorkLogger.Named("repo.product.Remove").With(zap.Uint64("productID", productID))

	if r.db == nil {
		l.Error("DB not initialized")
		return
	}

	_, err = r.db.Exec("DELETE FROM product WHERE id = $1", productID)
	if err != nil {
		l.Error("couldn't delete product")
		return
	}

	return
}
