package cart

import (
	"context"
	"database/sql"
	interfaces "github.com/Torebekov/shop-back/internal/interfaces/repository"
	"github.com/Torebekov/shop-back/internal/models"
	"github.com/Torebekov/shop-back/modules/logger"
	"go.uber.org/zap"
)

type cart struct {
	db  *sql.DB
	ctx context.Context
}

func New(db *sql.DB, ctx context.Context) interfaces.ICart {
	return &cart{
		db:  db,
		ctx: ctx,
	}
}

func (r *cart) Add(cartModel models.Cart) (err error) {
	l := logger.WorkLogger.Named("repo.cart.Add").With(zap.Any("cartModel", cartModel))

	if r.db == nil {
		l.Error("DB not initialized")
		return
	}

	_, err = r.db.Exec("INSERT INTO user_cart (user_id, product_id) VALUES ($1, $2)", cartModel.UserID, cartModel.ProductID)
	if err != nil {
		l.Error("couldn't insert cart")
		return
	}

	return
}

func (r *cart) Remove(cartModel models.Cart) (err error) {
	l := logger.WorkLogger.Named("repo.cart.Remove").With(zap.Any("cartModel", cartModel))

	if r.db == nil {
		l.Error("DB not initialized")
		return
	}

	_, err = r.db.Exec("DELETE FROM user_cart WHERE user_id = $1 AND product_id = $2", cartModel.UserID, cartModel.ProductID)
	if err != nil {
		l.Error("couldn't delete cart")
		return
	}

	return
}

func (r *cart) List(userID uint64) (products []models.Product, err error) {
	l := logger.WorkLogger.Named("repo.product.List")

	if r.db == nil {
		l.Error("DB not initialized")
		return
	}

	query := `
		SELECT
			product.id,
			product.name,
			category_id,
			image,
			price
		FROM
			product
		JOIN
			user_cart ON user_cart.product_id = product.ID AND user_id = $1`

	rows, err := r.db.Query(query, userID)
	if err != nil {
		l.Error("couldn't make request", zap.Error(err))
		return
	}
	defer rows.Close()

	for rows.Next() {
		var productModel models.Product
		err = rows.Scan(&productModel.ID, &productModel.Name, &productModel.CategoryID, &productModel.Image, &productModel.Price)
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
