package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shenAmit/gochat/handlers"
)

func ApiRoutes(app fiber.Router) {
	api := app.Group("/api")

	api.Get("/",handlers.ApiTest)

}
