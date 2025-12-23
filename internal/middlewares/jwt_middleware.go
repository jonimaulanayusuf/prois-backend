package middlewares

import (
	"strings"

	"prois-backend/internal/config"
	"prois-backend/internal/database"
	"prois-backend/internal/models"
	"prois-backend/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTProtected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")

		if authHeader == "" {
			return utils.ResBadRequest(c, "Missing token")
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			return []byte(config.GetEnv("JWT_SECRET", "secret")), nil
		})

		if err != nil || !token.Valid {
			return utils.ResInvalidCredential(c)
		}

		claims := token.Claims.(jwt.MapClaims)
		userID := uint(claims["user_id"].(float64))

		var user models.User
		if err := database.DB.First(&user, userID).Error; err != nil {
			return utils.ResInvalidCredential(c)
		}

		c.Locals("user_id", userID)
		c.Locals("username", claims["username"].(string))

		return c.Next()
	}
}
