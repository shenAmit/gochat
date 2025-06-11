package admin

import (
	"github.com/gofiber/fiber/v2"
)

func AdminDashboard(c *fiber.Ctx) error {
	return c.Render("/content/index", fiber.Map{
		"navigation": true,
		"footer":     true,
	}, "layouts/main.jet")
}
