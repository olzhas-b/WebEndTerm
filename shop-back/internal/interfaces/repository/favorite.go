package interfaces

import "github.com/Torebekov/shop-back/internal/models"

type IFavorite interface {
	Add(favoriteModel models.Favorite) (err error)
	Remove(favoriteModel models.Favorite) (err error)
	List(userID uint64) (products []models.Product, err error)
}
