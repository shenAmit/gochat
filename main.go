package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
	"github.com/shenAmit/gochat/config"
	"github.com/shenAmit/gochat/routes"
)

func main() {
	if _, err := os.Stat("./templates"); os.IsNotExist(err) {
		log.Fatal("Template directory './templates' does not exist")
	}

	engine := html.New("./templates", ".html")

	engine.Reload(true)

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./public")
	app.Static("/", "./static")

	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
	}

	config.ConnectMongoDB()
	routes.SetupRoutes(app)

	log.Println("Server starting on :3000")
	log.Fatal(app.Listen(":3000"))
}
