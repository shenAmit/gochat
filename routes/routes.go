package routes

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {
	admin := app.Group("/")

	AdminRoutes(admin)
	ApiRoutes(admin)
}