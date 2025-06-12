package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shenAmit/gochat/handlers/admin"
)

func AdminRoutes(app fiber.Router) {
	app.Get("/", admin.AdminDashboard)

}
