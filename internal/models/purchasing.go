package models

import "time"

type Purchasing struct {
	ID         uint               `json:"id" gorm:"primaryKey"`
	Date       time.Time          `json:"date"`
	SupplierID uint               `json:"supplier_id"`
	UserID     uint               `json:"user_id"`
	GrandTotal float64            `json:"grand_total"`
	Supplier   Supplier           `json:"supplier"`
	Details    []PurchasingDetail `json:"details"`
	CreatedAt  time.Time          `json:"created_at"`
}
