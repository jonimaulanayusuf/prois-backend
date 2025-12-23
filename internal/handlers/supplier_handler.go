package handlers

import (
	"prois-backend/internal/database"
	"prois-backend/internal/models"
	"prois-backend/internal/requests"
	"prois-backend/internal/resources"
	"prois-backend/internal/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func GetSuppliers(c *fiber.Ctx) error {
	var result []models.Supplier
	var total int64

	search := c.Query("search")
	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 10)

	if limit > 50 {
		limit = 50
	}

	offset := (page - 1) * limit
	query := database.DB.Model(&models.Supplier{})

	if search != "" {
		query = query.Where("name LIKE ?", "%"+search+"%")
	}

	query.Order("created_at DESC")
	query.Count(&total)
	query.Limit(limit).Offset(offset).Find(&result)

	list := []resources.SupplierResource{}
	for _, data := range result {
		list = append(list, resources.FromSupplier(data))
	}

	return utils.ResPagination(c, list, total, page, limit)
}

func CreateSupplier(c *fiber.Ctx) error {
	validate := validator.New()
	var data requests.CreateSupplierRequest

	if err := c.BodyParser(&data); err != nil {
		return utils.ResInvalidRequest(c)
	}

	if err := validate.Struct(&data); err != nil {
		return utils.ResValidationFailed(c, data, err)
	}

	row := models.Supplier{
		Name:    data.Name,
		Email:   data.Email,
		Address: data.Address,
	}

	database.DB.Create(&row)

	return utils.ResCreated(c, row)
}

func UpdateSupplier(c *fiber.Ctx) error {
	validate := validator.New()
	_id := c.Params("id")
	id := utils.DecryptID(_id)

	var row models.Supplier

	if err := database.DB.First(&row, id).Error; err != nil {
		return utils.ResNotFound(c)
	}

	var data requests.UpdateSupplierRequest
	if err := c.BodyParser(&data); err != nil {
		return utils.ResInvalidRequest(c)
	}

	if err := validate.Struct(&data); err != nil {
		return utils.ResValidationFailed(c, data, err)
	}

	if data.Name != nil {
		row.Name = *data.Name
	}
	if data.Email != nil {
		row.Email = *data.Email
	}
	if data.Address != nil {
		row.Address = *data.Address
	}

	database.DB.Save(&row)

	return utils.ResSuccess(c, resources.FromSupplier(row))
}

func DeleteSupplier(c *fiber.Ctx) error {
	_id := c.Params("id")
	id := utils.DecryptID(_id)

	var row models.Supplier

	if err := database.DB.First(&row, id).Error; err != nil {
		return utils.ResNotFound(c)
	}

	database.DB.Delete(&row)

	return utils.ResSuccess(c, nil)
}
