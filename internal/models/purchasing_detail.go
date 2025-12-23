package models

import "time"

type PurchasingDetail struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	PurchasingID string    `json:"purchasing_id" gorm:"type:varchar(191)"`
	ItemID       uint      `json:"item_id"`
	Qty          int       `json:"qty"`
	SubTotal     float64   `json:"sub_total"`
	Item         Item      `json:"item"`
	CreatedAt    time.Time `json:"created_at"`
}
