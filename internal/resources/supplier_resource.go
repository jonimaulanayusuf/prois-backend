package resources

import (
	"prois-backend/internal/models"
	"prois-backend/internal/utils"
)

type SupplierResource struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func FromSupplier(data models.Supplier) SupplierResource {
	return SupplierResource{
		ID:        utils.EncryptID(data.ID),
		Name:      data.Name,
		Email:     data.Email,
		Address:   data.Address,
		CreatedAt: data.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: data.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
