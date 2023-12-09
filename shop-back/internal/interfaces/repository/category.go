package interfaces

import (
	"github.com/Torebekov/shop-back/internal/models"
)

type ICategory interface {
	List() (categories []models.Category, err error)
}
