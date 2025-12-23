package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"prois-backend/internal/config"
	"prois-backend/internal/database"
	"prois-backend/internal/routes"
	"prois-backend/internal/utils"
)

func main() {
	config.LoadEnv()
	database.Connect()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return utils.ResMessage(c, fiber.StatusOK, "Welcome to Prois Backend 👋")
	})

	routes.AppRoutes(app)

	port := config.GetEnv("APP_PORT", "3001")
	log.Fatal(app.Listen(":" + port))
}
