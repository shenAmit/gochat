package admin

import (
	"github.com/gofiber/fiber/v2"
)

func AdminDashboard(c *fiber.Ctx) error {

	return c.Render("crud/index", fiber.Map{
		"title": "Admin Dashboard",
	}, "layouts/dashboard")
}
func loginPage(c *fiber.Ctx) error {
	return c.Render("content/authentication/sign-in", fiber.Map{
		"title": "Login - Chat",
	})
}
