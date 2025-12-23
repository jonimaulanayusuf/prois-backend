package routes

import (
	"prois-backend/internal/middlewares"

	"github.com/gofiber/fiber/v2"
)

func AppRoutes(app *fiber.App) {
	AuthRoutes(app.Group("/auth"))

	protected := app.Group("/", middlewares.JWTProtected())
	PurchasingRoutes(protected.Group("/purchasings"))
	ItemRoutes(protected.Group("/items"))
	SupplierRoutes(protected.Group("/suppliers"))
}
