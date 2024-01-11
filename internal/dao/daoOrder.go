package dao

import (
	"l0_ms/internal/models"
	"l0_ms/internal/repository"
	"l0_ms/internal/cache"
	"fmt"
	"log"
	"encoding/json"
	"gorm.io/gorm"
)


type Client struct {
	db              *gorm.DB
	cache           *cache.Cache
	orderRepository *repository.OrderRepository
}

func NewOrderClient(db *gorm.DB) *Client {
	return &Client{
		db:              db,
		cache:           cache.NewCache(),
		orderRepository: repository.NewOrderRepository(db),
	}
}

func (c *Client) Start() error {
	err := c.orderRepository.OrderAutoMigrate()
	if err != nil {
		return err
	}
	orders, err := c.orderRepository.GetAllOrder()
	if err != nil {
		return err
	}
	c.cache.AddOrders(orders)
	return nil
}

func (c *Client) GetOrder(orderUid string) (*models.Order, error) {
	order, err := c.cache.GetOrder(orderUid)
	if err != nil {
		log.Printf("order not found %v", err)
		return nil, fmt.Errorf("order not found")
	}

	return order, nil
}

func (c *Client) AddOrder(data []byte) error {
	var order models.Order
	err := json.Unmarshal(data, &order)

	if err != nil {
		log.Printf("error happened in JSON unmarshal. Err: %v", err)
	}

	fmt.Println("get Order UID: ", order.OrderUID)
	_, err = c.cache.GetOrder(order.OrderUID)
	if err == nil {
		log.Printf("such an order already exists. Err: %v", err)
		return fmt.Errorf("such an order already exists")
	}

	err = c.orderRepository.AddOrder(&order)
	if err != nil {
		log.Printf("Error create order: %v", err)
		return err
	}
	c.cache.AddOrder(order)
	return nil
}