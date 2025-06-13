package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func SetHTMXLayout(defaultLayout string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if c.Get("HX-Request") == "true" {
			c.Locals("layout", "") // no layout for HTMX
		} else {
			c.Locals("layout", defaultLayout)
		}
		return c.Next()
	}
}
