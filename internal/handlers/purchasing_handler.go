package handlers

import (
	"log"

	"prois-backend/internal/database"
	"prois-backend/internal/models"
	"prois-backend/internal/requests"
	"prois-backend/internal/resources"
	"prois-backend/internal/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetPurchasingHistory(c *fiber.Ctx) error {
	var result []models.Purchasing
	var total int64

	search := c.Query("search")
	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 10)

	if limit > 50 {
		limit = 50
	}

	offset := (page - 1) * limit
	query := database.DB.Model(&models.Purchasing{})

	if search != "" {
		query = query.Where("id LIKE ?", "%"+search+"%")
	}

	query.Preload("Supplier")
	query.Preload("Details.Item")
	query.Order("created_at DESC")
	query.Count(&total)
	query.Limit(limit).Offset(offset).Find(&result)

	list := []resources.PurchasingSummaryResource{}
	for _, data := range result {
		list = append(list, resources.FromPurchasingForSummary(data))
	}

	return utils.ResPagination(c, list, total, page, limit)
}

func GetPurchasingDetail(c *fiber.Ctx) error {
	id := c.Params("id")

	var row models.Purchasing

	err := database.DB.
		Preload("Supplier").
		Preload("Details.Item").
		Where("id = ?", id).
		First(&row).Error

	if err != nil {
		return utils.ResNotFound(c)
	}

	return utils.ResCreated(c, resources.FromPurchasing(row))
}

func CreatePurchasing(c *fiber.Ctx) error {
	validate := validator.New()
	userID := c.Locals("user_id").(uint)

	var data requests.CreatePurchasingRequest

	if err := c.BodyParser(&data); err != nil {
		return utils.ResInvalidRequest(c)
	}

	if err := validate.Struct(&data); err != nil {
		return utils.ResValidationFailed(c, data, err)
	}

	var grandTotal float64
	var purchasing models.Purchasing

	err := database.DB.Transaction(func(tx *gorm.DB) error {
		SupplierID := utils.DecryptID(data.SupplierID)
		var supplier models.Supplier

		if err := tx.First(&supplier, SupplierID).Error; err != nil {
			return fiber.NewError(fiber.StatusNotFound, "Supplier ID not found")
		}

		invoice, errInv := utils.GenerateInvoiceNumber()

		if errInv != nil {
			return errInv
		}

		purchasing = models.Purchasing{
			ID:         invoice,
			Date:       data.Date,
			SupplierID: *SupplierID,
			UserID:     userID,
		}

		if err := tx.Create(&purchasing).Error; err != nil {
			return err
		}

		for _, itemReq := range data.Items {
			ItemID := utils.DecryptID(itemReq.ItemID)
			var item models.Item

			if err := tx.First(&item, ItemID).Error; err != nil {
				return fiber.NewError(fiber.StatusNotFound, "Item not found")
			}

			if item.Stock < itemReq.Qty {
				return fiber.NewError(fiber.StatusBadRequest, "Stock not enough")
			}

			subTotal := float64(itemReq.Qty) * item.Price
			grandTotal += subTotal

			detail := models.PurchasingDetail{
				PurchasingID: purchasing.ID,
				ItemID:       item.ID,
				Qty:          itemReq.Qty,
				SubTotal:     subTotal,
			}

			if err := tx.Create(&detail).Error; err != nil {
				return fiber.NewError(fiber.StatusInternalServerError, "Failed create purchasing detail")
			}

			// update stock
			item.Stock -= itemReq.Qty
			if err := tx.Save(&item).Error; err != nil {
				return fiber.NewError(fiber.StatusInternalServerError, "Failed update stock")
			}
		}

		// update grand total
		if err := tx.Model(&purchasing).
			Update("grand_total", grandTotal).Error; err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Failed count grand total")
		}

		return nil
	})

	if err != nil {
		if fiberErr, ok := err.(*fiber.Error); ok {
			return utils.ResMessage(c, fiberErr.Code, fiberErr.Message)
		}

		return utils.ResInternalError(c, err.Error())
	}

	var row models.Purchasing

	errFind := database.DB.
		Preload("Supplier").
		Preload("Details.Item").
		Where("id = ?", purchasing.ID).
		First(&row).Error

	if errFind != nil {
		return utils.ResNotFound(c)
	}

	output := resources.FromPurchasing(row)

	go func() {
		if err := utils.SendPurchaseWebhook(output); err != nil {
			log.Println("Webhook error:", err)
		}
	}()

	return utils.ResCreated(c, output)
}
