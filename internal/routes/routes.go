package routes

import (
	"prois-backend/internal/middlewares"

	"github.com/gofiber/fiber/v2"
)

func AppRoutes(app *fiber.App) {
	AuthRoutes(app.Group("/auth"))
	InfoRoutes(app.Group("/info", middlewares.JWTProtected()))
	PurchasingRoutes(app.Group("/purchasings", middlewares.JWTProtected()))
	ItemRoutes(app.Group("/items", middlewares.JWTProtected()))
	SupplierRoutes(app.Group("/suppliers", middlewares.JWTProtected()))
}
