package routes

import (
	"prois-backend/internal/handlers"
	"prois-backend/internal/middlewares"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(route fiber.Router) {
	route.Post("/register", handlers.Register)
	route.Post("/login", handlers.Login)

	// protected
	route.Get("/me", middlewares.JWTProtected(), handlers.GetCurrentUser)
	route.Delete("/logout", middlewares.JWTProtected(), handlers.Logout)
}
