package main

import (
	"gin-pizza-order-tracker/infrastructure/database"
	"gin-pizza-order-tracker/internal/handlers"
	"gin-pizza-order-tracker/internal/repositories"
	"gin-pizza-order-tracker/internal/services"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		panic("DATABASE_URL is not set.")
	}

	db := database.NewPostgresDB(dsn)
	database.AutoMigrate(db) // development environment

	orderRepo := repositories.NewOrderRepository(db)
	orderService := services.NewOrderService(orderRepo)
	orderHandler := handlers.NewOrderHandler(orderService)

	r := gin.Default()
	orderHandler.RegisterRoutes(r)

	r.Run(":8080")
}
