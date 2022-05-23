package main

import (
	"github.com/RakibSiddiquee/go-fiber-jwt-auth/database"
	"github.com/RakibSiddiquee/go-fiber-jwt-auth/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()

	app := fiber.New()

	// Allow cors for cookie
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	app.Listen(":8000")
}
