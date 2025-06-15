package admin

import (
	"github.com/gofiber/fiber/v2"
)

func LoginPage(c *fiber.Ctx) error {
	return c.Render("content/authentication/sign-in", fiber.Map{
		"title": "Login - Chat",
	})
}
