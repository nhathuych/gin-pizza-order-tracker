package handlers

import (
	"gin-pizza-order-tracker/internal/models"
	"gin-pizza-order-tracker/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	service *services.OrderService
}

func NewOrderHandler(service *services.OrderService) *OrderHandler {
	return &OrderHandler{service: service}
}

func (h *OrderHandler) RegisterRoutes(r *gin.Engine) {
	orders := r.Group("/orders")
	{
		orders.GET("", h.getAll)
		orders.GET("/:id", h.getByID)
		orders.POST("", h.create)
		orders.PATCH("/:id/status", h.updateStatus)
		orders.DELETE("/:id", h.delete)
	}
}

func (h *OrderHandler) getAll(c *gin.Context) {
	orders, err := h.service.GetAllOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Failed to retrieve orders.")
		return
	}

	c.JSON(http.StatusOK, orders)
}

func (h *OrderHandler) getByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id."})
		return
	}

	order, err := h.service.GetOrder(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found."})
		return
	}

	c.JSON(http.StatusOK, order)
}

func (h *OrderHandler) create(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateOrder(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, order)
}

func (h *OrderHandler) updateStatus(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id."})
		return
	}

	newStatus := c.PostForm("status")
	if err := h.service.UpdateOrderStatus(uint(id), newStatus); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusSeeOther, "updated.")
}

func (h *OrderHandler) delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id."})
		return
	}

	if err := h.service.DeleteOrder(uint(id)); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusSeeOther, "deleted.")
}
