package handlers

import (
	"prois-backend/internal/database"
	"prois-backend/internal/models"
	"prois-backend/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func GetSummary(c *fiber.Ctx) error {
	var totalSuppliers int64
	var totalItems int64
	var totalPurchasings int64

	database.DB.Model(&models.Supplier{}).Count(&totalSuppliers)
	database.DB.Model(&models.Item{}).Count(&totalItems)
	database.DB.Model(&models.Purchasing{}).Count(&totalPurchasings)

	return utils.ResSuccess(c, fiber.Map{
		"total_items":       totalItems,
		"total_suppliers":   totalSuppliers,
		"total_purchasings": totalPurchasings,
	})
}
