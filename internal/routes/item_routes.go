package routes

import (
	"prois-backend/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func ItemRoutes(route fiber.Router) {
	route.Get("/", handlers.GetItems)
	route.Post("/", handlers.CreateItem)
	route.Patch("/:id", handlers.UpdateItem)
	route.Delete("/:id", handlers.DeleteItem)
}
