package database

import (
	"gin-pizza-order-tracker/internal/models"
	"log"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.Order{},
		&models.OrderItem{},
	)
	if err != nil {
		log.Fatalf("auto migrate failed: %v", err)
	}
}
