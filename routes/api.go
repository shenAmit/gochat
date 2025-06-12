package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shenAmit/gochat/handlers"
	// Import your API handlers here
	// "github.com/shenAmit/gochat/handlers/api"
)

func ApiRoutes(app fiber.Router) {
	app.Get("/", handlers.ApiTest)
}
