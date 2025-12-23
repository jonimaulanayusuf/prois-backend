package resources

import (
	"prois-backend/internal/models"
	"prois-backend/internal/utils"
	"time"
)

type PurchasingDetailResource struct {
	ID        string       `json:"id"`
	Item      ItemResource `json:"item"`
	Qty       int          `json:"qty"`
	SubTotal  float64      `json:"sub_total"`
	CreatedAt time.Time    `json:"created_at"`
}

func FromPurchasingDetail(data models.PurchasingDetail) PurchasingDetailResource {
	return PurchasingDetailResource{
		ID:        utils.EncryptID(data.ID),
		Item:      FromItem(data.Item),
		Qty:       data.Qty,
		SubTotal:  data.SubTotal,
		CreatedAt: data.CreatedAt,
	}
}
