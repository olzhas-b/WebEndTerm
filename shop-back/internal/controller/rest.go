package controller

import (
	"context"
	"github.com/Torebekov/shop-back/internal/clients/auth"
	core "github.com/Torebekov/shop-back/internal/core"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server struct {
	ctx  context.Context
	core *core.Core
	auth auth.Interface
}

func NewServer(
	ctx context.Context,
	core *core.Core,
	auth auth.Interface,
) *Server {
	return &Server{
		ctx:  ctx,
		core: core,
		auth: auth,
	}
}

func (s *Server) Run(port string) error {
	gin.SetMode(gin.DebugMode)
	router := gin.New()
	router.Use(gin.Recovery())

	v1 := router.Group("/api/v1")

	productRouter := v1.Group("product")
	{
		productRouter.GET("", s.AuthMiddleware(user, false), s.GetProducts)

		favoriteRouter := productRouter.Group("favorite")
		favoriteRouter.GET("", s.AuthMiddleware(user, true), s.GetFavorite)
		favoriteRouter.POST("/:id", s.AuthMiddleware(user, true), s.AddFavorite)
		favoriteRouter.DELETE("/:id", s.AuthMiddleware(user, true), s.RemoveFavorite)

		cartRouter := productRouter.Group("cart")
		cartRouter.GET("", s.AuthMiddleware(user, true), s.GetCart)
		cartRouter.POST("/:id", s.AuthMiddleware(user, true), s.AddCart)
		cartRouter.DELETE("/:id", s.AuthMiddleware(user, true), s.RemoveCart)
	}

	v1.GET("category", s.AuthMiddleware(user, false), s.GetCategories)

	userRouter := v1.Group("user")
	{
		userRouter.GET("")
		userRouter.PUT("")
	}

	return http.ListenAndServe(":"+port, router)
}
