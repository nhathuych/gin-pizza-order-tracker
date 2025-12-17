package repositories

import (
	"gin-pizza-order-tracker/internal/models"

	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

// constructor
func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) Create(order *models.Order) error {
	return r.db.Create(order).Error
}

func (r *OrderRepository) GetByID(id uint) (*models.Order, error) {
	var order models.Order
	err := r.db.Preload("Items").First(&order, id).Error
	return &order, err
}

func (r *OrderRepository) GetAll() ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Preload("Items").Order("created_at desc").Find(&orders).Error
	return orders, err
}

func (r *OrderRepository) UpdateStatus(id uint, status string) error {
	return r.db.Model(&models.Order{}).Where("id = ?", id).Update("status", status).Error
}

func (r *OrderRepository) Delete(id uint) error {
	return r.db.Delete(&models.Order{}, id).Error
}
