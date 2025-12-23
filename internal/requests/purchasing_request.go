package requests

type PurchasingItemRequest struct {
	ItemID string `json:"item_id"`
	Qty    int    `json:"qty"`
}

type CreatePurchasingRequest struct {
	SupplierID string                  `json:"supplier_id"`
	Items      []PurchasingItemRequest `json:"items"`
}
