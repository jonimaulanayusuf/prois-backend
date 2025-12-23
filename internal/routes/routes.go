package routes

import (
	"prois-backend/internal/middlewares"

	"github.com/gofiber/fiber/v2"
)

func AppRoutes(app *fiber.App) {
	guest := app.Group("/")
	protected := app.Group("/", middlewares.JWTProtected())

	AuthRoutes(guest.Group("/auth"))
	PurchasingRoutes(protected.Group("/purchasings"))
	ItemRoutes(protected.Group("/items"))
	SupplierRoutes(protected.Group("/suppliers"))
}
