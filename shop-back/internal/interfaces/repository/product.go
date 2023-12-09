package interfaces

import (
	"github.com/Torebekov/shop-back/internal/models"
)

type IProduct interface {
	List(searchText string, categoryID, userID uint64) (products []models.ProductDTO, err error)
	ByID(productID uint64) (productModel models.Product, err error)
	Add(productModel models.Product) (err error)
	Remove(productID uint64) (err error)
}
