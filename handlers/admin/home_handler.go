package admin

import (
	"github.com/gofiber/fiber/v2"
)

func AdminDashboard(c *fiber.Ctx) error {
	return c.Render("content/index", fiber.Map{
		"title":      "Admin Dashboard",
		"navigation": true,
		"footer":     true,
	})
}
