package resources

import (
	"prois-backend/internal/models"
	"prois-backend/internal/utils"
)

type ItemResource struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Stock     int     `json:"stock"`
	Price     float64 `json:"price"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

func FromItem(data models.Item) ItemResource {
	return ItemResource{
		ID:        utils.EncryptID(data.ID),
		Name:      data.Name,
		Stock:     data.Stock,
		Price:     data.Price,
		CreatedAt: data.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: data.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
