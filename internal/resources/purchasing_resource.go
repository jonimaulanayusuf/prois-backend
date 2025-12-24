package resources

import (
	"prois-backend/internal/models"
)

type PurchasingResource struct {
	ID         string                     `json:"id"`
	Date       string                     `json:"date"`
	Supplier   SupplierResource           `json:"supplier"`
	GrandTotal float64                    `json:"grand_total"`
	Details    []PurchasingDetailResource `json:"details"`
	CreatedAt  string                     `json:"created_at"`
}

func FromPurchasing(data models.Purchasing) PurchasingResource {
	var details []PurchasingDetailResource
	for _, data := range data.Details {
		details = append(details, FromPurchasingDetail(data))
	}

	return PurchasingResource{
		ID:         data.ID,
		Date:       data.Date,
		Supplier:   FromSupplier(data.Supplier),
		GrandTotal: data.GrandTotal,
		Details:    details,
		CreatedAt:  data.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}
