package models

import (
	"gorm.io/gorm"
)

type OrderItem struct {
	gorm.Model
	OrderID      uint   `gorm:"index;not null" json:"order_id"`
	Size         string `gorm:"not null" json:"size"`
	Pizza        string `gorm:"not null" json:"pizza"`
	Instructions string `json:"instructions"`
}
