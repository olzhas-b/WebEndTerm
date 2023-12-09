package favorite

import (
	"context"
	"database/sql"
	interfaces "github.com/Torebekov/shop-back/internal/interfaces/repository"
	"github.com/Torebekov/shop-back/internal/models"
	"github.com/Torebekov/shop-back/modules/logger"
	"go.uber.org/zap"
)

type favorite struct {
	db  *sql.DB
	ctx context.Context
}

func New(db *sql.DB, ctx context.Context) interfaces.IFavorite {
	return &favorite{
		db:  db,
		ctx: ctx,
	}
}

func (r *favorite) Add(favoriteModel models.Favorite) (err error) {
	l := logger.WorkLogger.Named("repo.favorite.AddFavorite").With(zap.Any("favoriteModel", favoriteModel))

	if r.db == nil {
		l.Error("DB not initialized")
		return
	}

	_, err = r.db.Exec("INSERT INTO user_favorite (user_id, product_id) VALUES ($1, $2)", favoriteModel.UserID, favoriteModel.ProductID)
	if err != nil {
		l.Error("couldn't insert favorite")
		return
	}

	return
}

func (r *favorite) Remove(favoriteModel models.Favorite) (err error) {
	l := logger.WorkLogger.Named("repo.favorite.AddFavorite").With(zap.Any("favoriteModel", favoriteModel))

	if r.db == nil {
		l.Error("DB not initialized")
		return
	}

	_, err = r.db.Exec("DELETE FROM user_favorite WHERE user_id = $1 AND product_id = $2", favoriteModel.UserID, favoriteModel.ProductID)
	if err != nil {
		l.Error("couldn't delete favorite")
		return
	}

	return
}

func (r *favorite) List(userID uint64) (products []models.Product, err error) {
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
			user_favorite ON user_favorite.product_id = product.ID AND user_id = $1`

	rows, err := r.db.Query(query, userID)
	if err != nil {
		l.Error("couldn't make request", zap.Error(err))
		return
	}
	defer rows.Close()

	rows.Next()
	{
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
