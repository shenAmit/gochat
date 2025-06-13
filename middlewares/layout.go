package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func HTMXCheck() fiber.Handler {
	return func(c *fiber.Ctx) error {
		isHTMX := c.Get("HX-Request") == "true"
		c.Locals("isHTMX", isHTMX) // Store the boolean in the context
		return c.Next()            // Pass control to the next handler
	}
}
