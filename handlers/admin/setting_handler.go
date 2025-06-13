package admin

import "github.com/gofiber/fiber/v2"

func isHTMX(c *fiber.Ctx) bool {
	return c.Get("HX-Request") == "true"
}

func SettingHandler(c *fiber.Ctx) error {
	if isHTMX(c) {
		return c.Render("crud/setting", fiber.Map{
			"title": "Setting",
		})
	}

	return c.Render("crud/setting", fiber.Map{
		"title": "Setting",
	}, "layouts/dashboard")
}
