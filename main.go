package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/jet/v2"
	"github.com/shenAmit/gochat/config"
	"github.com/shenAmit/gochat/routes"
)

func main() {
	engin := jet.New("./views", ".jet")
	app := fiber.New(fiber.Config{
		Views: engin,
	})
	app.Static("/", "./public")


	config.ConnectMongoDB()
	routes.SetupRoutes(app)


	app.Listen(":3000")
}
