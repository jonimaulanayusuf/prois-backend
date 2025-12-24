package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"prois-backend/internal/config"
	"prois-backend/internal/database"
	"prois-backend/internal/routes"
	"prois-backend/internal/utils"
)

func main() {
	config.LoadEnv()
	database.Connect()

	app := fiber.New()

	allowedOrigins := config.GetEnv("ALLOWED_ORIGINS", "*") // ini ngambil dari env dan sudah di make sure
	app.Use(cors.New(cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowMethods:     "GET,POST,PUT,PATCH,DELETE,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return utils.ResMessage(c, fiber.StatusOK, "Welcome to Prois Backend 👋")
	})

	routes.AppRoutes(app)

	port := config.GetEnv("APP_PORT", "3001")
	log.Fatal(app.Listen(":" + port))
}
