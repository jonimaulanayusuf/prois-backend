package handlers

import (
	"prois-backend/internal/config"
	"prois-backend/internal/database"
	"prois-backend/internal/models"
	"prois-backend/internal/requests"
	"prois-backend/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var data requests.RegisterRequest
	if err := c.BodyParser(&data); err != nil {
		return utils.ResInvalidRequest(c)
	}

	if data.Username == "" || data.Password == "" {
		return utils.ResBadRequest(c, "Username & password required")
	}

	if data.Password != data.PasswordConfirmation {
		return utils.ResBadRequest(c, "Confirm your account password")
	}

	hashedPassword, err := utils.HashPassword(data.Password)
	if err != nil {
		return utils.ResInternalError(c, "Failed to hash password")
	}

	user := models.User{
		Username: data.Username,
		Password: hashedPassword,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return utils.ResBadRequest(c, "Username already exists")
	}

	token, err := utils.GenerateJWT(
		user.ID,
		user.Username,
		config.GetEnv("JWT_SECRET", "secret"),
	)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to generate token")
	}

	return utils.ResCreated(c, fiber.Map{
		"user":  user,
		"token": token,
	})
}

func Login(c *fiber.Ctx) error {
	var data requests.LoginRequest
	if err := c.BodyParser(&data); err != nil {
		return utils.ResInvalidRequest(c)
	}

	var user models.User
	if err := database.DB.Where("username = ?", data.Username).First(&user).Error; err != nil {
		return utils.ResInvalidCredential(c)
	}

	if !utils.CheckPasswordHash(data.Password, user.Password) {
		return utils.ResInvalidCredential(c)
	}

	token, err := utils.GenerateJWT(
		user.ID,
		user.Username,
		config.GetEnv("JWT_SECRET", "secret"),
	)

	if err != nil {
		return utils.ResInternalError(c, "Failed to generate token")
	}

	return utils.ResSuccess(c, fiber.Map{
		"user":  user,
		"token": token,
	})
}

func GetCurrentUser(c *fiber.Ctx) error {
	userID := c.Locals("user_id")

	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		return utils.ResNotFound(c)
	}

	return utils.ResSuccess(c, fiber.Map{
		"user": user,
	})
}
