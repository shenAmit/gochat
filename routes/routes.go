package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func SetupRoutes(app *fiber.App) {
	app.Use(logger.New())
	app.Use(cors.New())
	app.Use(recover.New())

	adminGroup := app.Group("/")
	AdminRoutes(adminGroup)

	apiGroup := app.Group("/api")
	ApiRoutes(apiGroup)
}
