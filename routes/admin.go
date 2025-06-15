package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shenAmit/gochat/handlers/admin"
	"github.com/shenAmit/gochat/middleware"
)

func AdminRoutes(app fiber.Router) {
	app.Get("/login", admin.LoginPage)
	app.Post("/login", admin.LoginHandler)
	app.Get("/debug/session", admin.DebugSession)
	app.Use(middleware.IsAdmin)

	app.Get("/logout", admin.LogoutHandler)
	app.Get("/", admin.AdminDashboard)
	app.Get("/setting", admin.SettingHandler)
}
