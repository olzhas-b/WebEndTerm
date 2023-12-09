package interfaces

import (
	"database/sql"
	"github.com/Torebekov/shop-back/internal/models"
)

type IProduct interface {
	GetDB() *sql.DB
	List(searchText string, categoryID uint64) (products []models.Product, err error)
}
