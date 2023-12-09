package tools

import (
	"github.com/gofiber/fiber/v2"
	"strings"
)

func GetToken(ctx *fiber.Ctx) (token string) {
	header := ctx.Get("Authorization")
	if header == "" {
		return
	}
	parsedHeader := strings.Split(header, " ")
	if len(parsedHeader) != 2 || parsedHeader[0] != "Bearer" {
		return
	}
	token = parsedHeader[1]
	return
}
