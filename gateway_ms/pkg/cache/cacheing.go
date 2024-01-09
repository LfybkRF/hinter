package cache

import (
	"gateway_l0_ms/internal/models"
	"fmt"
	"log"
	"errors"
)

type Cache struct {
	cache map[string] models.Order
}

func NewCache() *Cache {
	return &Cache{cache: make(map[string] models.Order)}
}

func (c *Cache) GetOrder(orderUid string) (*models.Order, error) {
	order, ok := c.cache[orderUid]
	if ok {
		return &order, nil
	}
	log.Println("UUID: " + order.OrderUID + "not found")
	return nil, fmt.Errorf("order not found")
}

func (c *Cache) AddOrders(orders []models.Order) error {
	for _, order := range orders {
		c.AddOrder(order)
	}
	return nil
}

func (c *Cache) AddOrder(order models.Order) error {
	if _, ok := c.cache[order.OrderUID]; ok {
		log.Println("UUID: " + order.OrderUID + "already exists")
		return errors.New("already exists")
	}
	c.cache[order.OrderUID] = order
	return nil
}