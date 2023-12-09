package services

import (
	"context"
	"database/sql"
	"github.com/Torebekov/shop-back/config"
	irepo "github.com/Torebekov/shop-back/internal/interfaces/repository"
	"github.com/Torebekov/shop-back/internal/repository/cart"
	"github.com/Torebekov/shop-back/internal/repository/category"
	"github.com/Torebekov/shop-back/internal/repository/favorite"
	"github.com/Torebekov/shop-back/internal/repository/product"
)

type Core struct {
	ctx     context.Context
	cfg     *config.Config
	sqlDB   *sql.DB
	isDebug bool

	productRepo  irepo.IProduct
	categoryRepo irepo.ICategory
	favoriteRepo irepo.IFavorite
	cartRepo     irepo.ICart
}

func New(
	ctx context.Context,
	cfg *config.Config,
	sqlDB *sql.DB,
) *Core {
	core := &Core{
		ctx:   ctx,
		cfg:   cfg,
		sqlDB: sqlDB,
	}

	return core.InitRepositories()
}

func (s *Core) InitRepositories() *Core {
	return s.
		InitProduct().
		InitCategory().
		InitFavorite().
		InitCart()
}

func (s *Core) InitProduct() *Core {
	s.productRepo = product.New(s.sqlDB, s.ctx)
	return s
}

func (s *Core) InitCategory() *Core {
	s.categoryRepo = category.New(s.sqlDB, s.ctx)
	return s
}

func (s *Core) InitCart() *Core {
	s.cartRepo = cart.New(s.sqlDB, s.ctx)
	return s
}

func (s *Core) InitFavorite() *Core {
	s.favoriteRepo = favorite.New(s.sqlDB, s.ctx)
	return s
}

func (s *Core) Product() irepo.IProduct {
	return s.productRepo
}

func (s *Core) Category() irepo.ICategory {
	return s.categoryRepo
}

func (s *Core) Favorite() irepo.IFavorite {
	return s.favoriteRepo
}

func (s *Core) Cart() irepo.ICart {
	return s.cartRepo
}

func (s *Core) Config() *config.Config {
	return s.cfg
}
