package filter

import "github.com/gofiber/fiber/v2"

type User struct {
	Page       int
	Size       int
	Sort       string
	Order      string
	SearchText string
}

func (u *User) Fill(ctx *fiber.Ctx) {

}
