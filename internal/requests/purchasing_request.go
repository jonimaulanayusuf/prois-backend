package requests

type PurchasingItemRequest struct {
	ItemID string `json:"item_id" validate:"required"`
	Qty    int    `json:"qty" validate:"required,min=1"`
}

type CreatePurchasingRequest struct {
	SupplierID string                  `json:"supplier_id" validate:"required"`
	Date       string                  `json:"date" validate:"required,datetime=2006-01-02"`
	Items      []PurchasingItemRequest `json:"items" validate:"required,min=1,dive"`
}
