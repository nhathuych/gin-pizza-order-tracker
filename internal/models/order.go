package models

import (
	"gorm.io/gorm"
)

var (
	OrderStatuses = []string{"Order placed", "Preparing", "Baking", "Quality Check", "Ready"}

	PizzaTypes = []string{
		"Margherita",
		"Pepperoni",
		"Vegetarian",
		"Hawaiian",
		"Bbq Chicken",
		"Meat Lovers",
		"Buffalo Chicken",
		"Supreme",
		"Truffle Mushroom",
		"Four Cheese",
	}

	PizzaSizes = []string{
		"Small", "Medium", "Large", "X-Large",
	}
)

type Order struct {
	gorm.Model
	Status       string      `gorm:"not null" json:"status"`
	CustomerName string      `gorm:"not null" json:"customer_name"`
	Phone        string      `gorm:"not null" json:"phone"`
	Address      string      `gorm:"not null" json:"address"`
	Items        []OrderItem `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE" json:"pizzas"`
}
