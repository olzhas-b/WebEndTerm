package product

import (
	"context"
	"database/sql"
	"fmt"
	interfaces "github.com/Torebekov/shop-back/internal/interfaces/repository"
	"github.com/Torebekov/shop-back/internal/models"
	"github.com/Torebekov/shop-back/modules/logger"
	"go.uber.org/zap"
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
	l := logger.WorkLogger.Named("repo.product.List").With(zap.String("searchText", searchText), zap.Uint64("categoryID", categoryID))

	if r.db == nil {
		l.Error("DB not initialized")
		return
	}

	query := fmt.Sprintf(`
		SELECT
			product.id,
			product.name,
			category.name,
			image,
			price,
			CASE WHEN user_favorite.product_id IS NULL THEN FALSE ELSE TRUE END
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
	addWhereClause(" category_id = %d", categoryID)

	rows, err := r.db.Query(query)
	if err != nil {
		l.Error("couldn't make request", zap.Error(err))
		return
	}
	defer rows.Close()

	rows.Next()
	{
		var productModel models.ProductDTO
		err = rows.Scan(&productModel.ID, &productModel.Name, &productModel.CategoryName, &productModel.Image, &productModel.Price, &productModel.IsFavorite)
		if err != nil {
			l.Error("couldn't scan product", zap.Error(err))
			return
		}

		products = append(products, productModel)
	}

	if err = rows.Err(); err != nil {
		l.Error("iteration error", zap.Error(err))
		return
	}

	return
}
