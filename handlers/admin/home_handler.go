package admin

import (
	"github.com/gofiber/fiber/v2"
)

func AdminDashboard(c *fiber.Ctx) error {
	val := c.Locals("isHTMX")
	isHTMX, ok := val.(bool)
	if !ok {
		isHTMX = false
	}

	if isHTMX {
		return c.Render("crud/index", fiber.Map{
			"title": "Admin Dashboard",
		})
	}
	return c.Render("layouts/dashboard", fiber.Map{
		"title": "Admin Dashboard",
	}, "layouts/dashboard")
}
