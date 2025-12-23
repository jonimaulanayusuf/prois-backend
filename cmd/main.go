package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"prois-backend/internal/config"
	"prois-backend/internal/database"
	"prois-backend/internal/routes"
)

func main() {
	config.LoadEnv()
	database.Connect()

	app := fiber.New()

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	routes.AppRoutes(app)

	port := config.GetEnv("APP_PORT", "3000")
	log.Fatal(app.Listen(":" + port))
}
