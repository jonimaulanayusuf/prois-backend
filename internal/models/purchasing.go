package models

import "time"

type Purchasing struct {
	ID         string             `json:"id" gorm:"type:varchar(191);primaryKey"`
	Date       string             `json:"date"`
	SupplierID uint               `json:"supplier_id"`
	UserID     uint               `json:"user_id"`
	GrandTotal float64            `json:"grand_total"`
	Supplier   Supplier           `json:"supplier"`
	Details    []PurchasingDetail `json:"details"`
	CreatedAt  time.Time          `json:"created_at"`
}
