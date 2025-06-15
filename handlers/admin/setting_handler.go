package admin

import "github.com/gofiber/fiber/v2"

func SettingHandler(c *fiber.Ctx) error {

	return c.Render("crud/setting", fiber.Map{
		"title": "Setting",
	}, "layouts/dashboard")
}
