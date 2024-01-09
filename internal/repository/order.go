package repository

import (
	"l0_ms/internal/models"
	"fmt"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (r *OrderRepository) OrderAutoMigrate() error {
	err := r.db.AutoMigrate(&models.Order{})
	if err != nil {
		return fmt.Errorf("error AutoMigrate: %v", err)
	}
	return nil
}

func (r *OrderRepository) AddOrder(order *models.Order) error {
	tx := r.db.Create(&order)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (r *OrderRepository) GetAllOrder() ([]models.Order, error) {
	var orders []models.Order
	dbOrders := r.db.Find(&orders)

	if dbOrders.Error != nil {
		return nil, dbOrders.Error
	}

	return orders, nil
}