package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shenAmit/gochat/handlers"
)

func ApiRoutes(app fiber.Router) {
	app.Get("/", handlers.ApiTest)
	app.Post("/register", handlers.Register)

}
