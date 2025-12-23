package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ResInvalidCredential(c *fiber.Ctx) error {
	return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
		"message": "Invalid credential",
	})
}

func ResPagination(c *fiber.Ctx, data interface{}, total int64, page int, limit int) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success",
		"result": fiber.Map{
			"data":        data,
			"total":       total,
			"page":        page,
			"limit":       limit,
			"total_pages": (total + int64(limit) - 1) / int64(limit),
		},
	})
}

func ResMessage(c *fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(fiber.Map{
		"message": message,
	})
}

func ResNotFound(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"message": "Not found",
	})
}

func ResInvalidRequest(c *fiber.Ctx) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"message": "Invalid request body",
	})
}

func ResBadRequest(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"message": message,
	})
}

func ResInternalError(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"message": message,
	})
}

func ResMissingToken(c *fiber.Ctx) error {
	return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
		"message": "Missing token",
	})
}

func ResCreated(c *fiber.Ctx, data interface{}) error {
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Created",
		"data":    data,
	})
}

func ResSuccess(c *fiber.Ctx, data interface{}) error {
	resp := fiber.Map{
		"message": "Ok",
	}

	if data != nil {
		resp["data"] = data
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}

func ResValidationFailed(c *fiber.Ctx, data interface{}, err error) error {
	errors := make(map[string]string)
	for _, e := range err.(validator.ValidationErrors) {
		jsonField := GetJSONFieldName(data, e.Field())
		switch e.Tag() {
		case "required":
			errors[jsonField] = "This field is required"
		case "max":
			errors[jsonField] = "Maximum length is " + e.Param()
		case "gte":
			errors[jsonField] = "Must be greater than or equal to " + e.Param()
		case "email":
			errors[jsonField] = "Invalid email address"
		}
	}

	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"message": "Validation failed",
		"errors":  errors,
	})
}
