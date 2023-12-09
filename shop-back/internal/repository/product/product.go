package product

import (
	"context"
	"database/sql"
	"github.com/Torebekov/shop-back/internal/models"
	"github.com/Torebekov/shop-back/modules/logger"
)

type product struct {
	db  *sql.DB
	ctx context.Context
}

func New(db *sql.DB, ctx context.Context) *product {
	return &product{
		db:  db,
		ctx: ctx,
	}
}

func (r *product) GetDB() *sql.DB {
	return r.db
}

func (r *product) List(string, uint64) (products []models.Product, err error) {
	l := logger.WorkLogger.Named("repo.product.List")

	if r.db == nil {
		l.Error("DB not initialized")
		return
	}

	return
}
