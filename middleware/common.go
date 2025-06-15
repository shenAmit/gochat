package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func IsAdmin(c *fiber.Ctx) error {
	admin := c.Cookies("session")
	if admin == "" {
		return c.Redirect("/login")
	}
	return c.Next()
}
