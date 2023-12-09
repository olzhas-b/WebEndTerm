package category

import (
	"context"
	"database/sql"
	interfaces "github.com/Torebekov/shop-back/internal/interfaces/repository"
	"github.com/Torebekov/shop-back/internal/models"
	"github.com/Torebekov/shop-back/modules/logger"
	"go.uber.org/zap"
)

type category struct {
	db  *sql.DB
	ctx context.Context
}

func New(db *sql.DB, ctx context.Context) interfaces.ICategory {
	return &category{
		db:  db,
		ctx: ctx,
	}
}

func (r *category) List() (categories []models.Category, err error) {
	l := logger.WorkLogger.Named("repo.category.List")

	if r.db == nil {
		l.Error("DB not initialized")
		return
	}

	rows, err := r.db.Query(`SELECT id, name FROM category`)
	if err != nil {
		l.Error("couldn't make request", zap.Error(err))
		return
	}
	defer rows.Close()

	rows.Next()
	{
		var categoryModel models.Category
		err = rows.Scan(&categoryModel.ID, &categoryModel.Name)
		if err != nil {
			l.Error("couldn't scan category", zap.Error(err))
			return
		}

		categories = append(categories, categoryModel)
	}

	if err = rows.Err(); err != nil {
		l.Error("iteration error", zap.Error(err))
		return
	}

	return
}
