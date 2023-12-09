package interfaces

import (
	"github.com/Torebekov/shop-back/internal/models"
)

type IProduct interface {
	List(searchText string, categoryID, userID uint64) (products []models.ProductDTO, err error)
}
