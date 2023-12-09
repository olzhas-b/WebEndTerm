package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/olzhas-b/PetService/authService/pkg/services"
)

type Handler struct {
	services *services.Services
}

func NewHandler(services *services.Services) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitializeRoutes(srv *fiber.App) error {
	srv.Use(
		recover.New(),
	)
	h.AddRoutes(srv)
	return nil
}

func (h *Handler) AddRoutes(srv *fiber.App) {

	v1 := srv.Group("/v1/auth")

	v1.Post("/sign-in", h.CtlSignIn)
	v1.Post("/sign-up", h.CtlCreateUser)
	v1.Post("/sign-out", h.CtlSignOut)
	v1.Post("/refresh", h.CtlRefreshToken)
	v1.Post("/check", h.CtlCheck)

}
