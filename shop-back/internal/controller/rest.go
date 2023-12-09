package controller

import (
	"context"
	core "github.com/Torebekov/shop-back/internal/core"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server struct {
	ctx  context.Context
	core *core.Core
}

func NewServer(
	ctx context.Context,
	core *core.Core,
) *Server {
	return &Server{
		ctx:  ctx,
		core: core,
	}
}

func (s *Server) Run(port string) error {
	gin.SetMode(gin.DebugMode)
	router := gin.New()
	router.Use(gin.Recovery())

	v1 := router.Group("/api/v1")

	productRouter := v1.Group("product")
	{
		productRouter.GET("", s.GetProducts)

		favoriteRouter := productRouter.Group("favorite")
		favoriteRouter.GET("")
		favoriteRouter.POST("")
		favoriteRouter.DELETE("")

		cartRouter := productRouter.Group("cart")
		cartRouter.GET("")
		cartRouter.POST("")
		cartRouter.DELETE("")
	}

	v1.GET("category")

	userRouter := v1.Group("user")
	{
		userRouter.GET("")
		userRouter.PUT("")
	}

	return http.ListenAndServe(":"+port, router)
}
