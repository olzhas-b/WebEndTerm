package common

import (
	"github.com/gofiber/fiber/v2"
	"github.com/olzhas-b/PetService/authService/consts"
	"net/http"
)

type response struct {
	Version string      `json:"version,omitempty"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

func GenShortResponse(ctx *fiber.Ctx, code int, data interface{}, message string) error {
	status := getStatusFromCode(code)
	if data != nil {
		return ctx.Status(status).JSON(data)
	}
	return ctx.SendStatus(status)
}

func getStatusFromCode(code int) int {
	switch code {
	case consts.Success:
		return http.StatusOK
	default:
		return http.StatusInternalServerError
	}
}

func GenResponse(c *fiber.Ctx, code int, data interface{}, message string) error {
	status := getStatusFromCode(code)
	response := response{consts.VERSION, code, data, message}

	return c.Status(status).JSON(response)
}
