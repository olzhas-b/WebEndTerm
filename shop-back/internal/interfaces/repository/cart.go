package interfaces

import "github.com/Torebekov/shop-back/internal/models"

type ICart interface {
	Add(cartModel models.Cart) (err error)
	Remove(cartModel models.Cart) (err error)
	List(userID uint64) (products []models.Product, err error)
}
