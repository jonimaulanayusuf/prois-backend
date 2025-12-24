package routes

import (
	"prois-backend/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func InfoRoutes(route fiber.Router) {
	route.Get("/", handlers.GetSummary)
}
