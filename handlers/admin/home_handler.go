package admin

import (
	"github.com/gofiber/fiber/v2"
)

func AdminDashboard(c *fiber.Ctx) error {
	layout, _ := c.Locals("layout").(string)
	return c.Render("crud/index", fiber.Map{
		"title": "Admin Dashboard",
	}, layout)
}
