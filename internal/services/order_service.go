package services

import (
	"errors"
	"gin-pizza-order-tracker/internal/models"
	"gin-pizza-order-tracker/internal/repositories"
)

type OrderService struct {
	repo *repositories.OrderRepository
}

func NewOrderService(repo *repositories.OrderRepository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) CreateOrder(order *models.Order) error {
	if order.CustomerName == "" {
		return errors.New("customer name is required")
	}

	order.Status = models.OrderStatuses[0] // default value
	return s.repo.Create(order)
}

func (s *OrderService) GetOrder(id uint) (*models.Order, error) {
	return s.repo.GetByID(id)
}

func (s *OrderService) GetAllOrders() ([]models.Order, error) {
	return s.repo.GetAll()
}

func (s *OrderService) UpdateOrderStatus(id uint, status string) error {
	return s.repo.UpdateStatus(id, status)
}

func (s *OrderService) DeleteOrder(id uint) error {
	return s.repo.Delete(id)
}
