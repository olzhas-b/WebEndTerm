package transport

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/olzhas-b/PetService/authService/config"
	"github.com/olzhas-b/PetService/authService/pkg/transport/restful/handler"
)

type Server struct {
	HTTPServer *fiber.App
}

func (s *Server) RunHTTPServer(config *config.Config, routes *handler.Handler) error {
	s.HTTPServer = fiber.New(fiber.Config{
		DisableStartupMessage: true,
		ReduceMemoryUsage:     true,
		Immutable:             true,
	})
	//s.HTTPServer.Use(cache.New(cache.Config{
	//	Next: func(c *fiber.Ctx) bool {
	//		return strings.Contains(c.Route().Path, "/ws")
	//	},
	//}))
	//s.HTTPServer.Use(cors.New(cors.Config{
	//	AllowOrigins:     strings.Join(config.Cors.AllowOrigins, ","),
	//	AllowHeaders:     strings.Join(config.Cors.AllowHeaders, ","),
	//	AllowCredentials: config.Cors.AllowCredentials,
	//	AllowMethods:     strings.Join(config.Cors.AllowMethods, ","),
	//	ExposeHeaders:    strings.Join(config.Cors.ExposeHeaders, ","),
	//	MaxAge:           int((config.Cors.MaxAge * time.Hour).Seconds()),
	//}))
	//
	//p := fiberprometheus.New("pe4kin")
	//p.RegisterAt(s.HTTPServer, "/metrics")
	//s.HTTPServer.Use(p.Middleware)
	//s.HTTPServer.Use(recover.New())

	err := routes.InitializeRoutes(s.HTTPServer)
	if err != nil {
		return err
	}
	return s.HTTPServer.Listen("0.0.0.0:8081")
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.HTTPServer.Shutdown()
}
