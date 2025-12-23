package routes

import (
	"prois-backend/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func SupplierRoutes(route fiber.Router) {
	route.Get("/", handlers.GetSuppliers)
	route.Post("/", handlers.CreateSupplier)
	route.Patch("/:id", handlers.UpdateSupplier)
	route.Delete("/:id", handlers.DeleteSupplier)
}
