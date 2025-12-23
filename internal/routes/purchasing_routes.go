package routes

import (
	"prois-backend/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func PurchasingRoutes(route fiber.Router) {
	route.Post("/", handlers.CreatePurchasing)
	route.Get("/", handlers.GetPurchasingHistory)
	route.Get("/:id", handlers.GetPurchasingDetail)
}
