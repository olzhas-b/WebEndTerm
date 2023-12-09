package interfaces

import (
	"github.com/Torebekov/shop-back/config"
	irepo "github.com/Torebekov/shop-back/internal/interfaces/repository"
)

type ICore interface {
	Config() *config.Config
	SmsConfirmationRepo() irepo.IProduct
}
