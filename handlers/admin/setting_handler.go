package admin

import "github.com/gofiber/fiber/v2"

func SettingHandler(c *fiber.Ctx) error {
	layout, _ := c.Locals("layout").(string)
	return c.Render("crud/setting", fiber.Map{
		"title": "Setting",
	}, layout)
}
