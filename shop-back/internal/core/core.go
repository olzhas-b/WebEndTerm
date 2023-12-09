package services

import (
	"context"
	"database/sql"
	"github.com/Torebekov/shop-back/config"
	irepo "github.com/Torebekov/shop-back/internal/interfaces/repository"
	"github.com/Torebekov/shop-back/internal/repository/product"
)

type Core struct {
	ctx     context.Context
	cfg     *config.Config
	sqlDB   *sql.DB
	isDebug bool

	productRepo irepo.IProduct
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
		InitProduct()
}

func (s *Core) InitProduct() *Core {
	s.productRepo = product.New(s.sqlDB, s.ctx)
	return s
}

func (s *Core) Product() irepo.IProduct {
	return s.productRepo
}

func (s *Core) Config() *config.Config {
	return s.cfg
}
