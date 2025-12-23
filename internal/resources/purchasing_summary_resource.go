package resources

import (
	"prois-backend/internal/models"
	"prois-backend/internal/utils"
	"time"
)

type PurchasingSummaryResource struct {
	ID           string    `json:"id"`
	Date         time.Time `json:"date"`
	SupplierName string    `json:"supplier_name"`
	GrandTotal   float64   `json:"grand_total"`
	CreatedAt    string    `json:"created_at"`
}

func FromPurchasingForSummary(data models.Purchasing) PurchasingSummaryResource {
	var details []PurchasingDetailResource
	for _, data := range data.Details {
		details = append(details, FromPurchasingDetail(data))
	}

	return PurchasingSummaryResource{
		ID:           utils.EncryptID(data.ID),
		Date:         data.Date,
		SupplierName: data.Supplier.Name,
		GrandTotal:   data.GrandTotal,
		CreatedAt:    data.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}
