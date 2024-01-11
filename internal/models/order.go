package models

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type Order struct {
	OrderUID          string    `json:"order_uid" gorm:"primaryKey"`
	TrackNumber       string    `json:"track_number"`
	Entry             string    `json:"entry"`
	Delivery          Delivery  `json:"delivery" serializer:"json"`
	Payment           Payment   `json:"payment" serializer:"json"`
	Items             Items     `json:"items" serializer:"json"`
	Locale            string    `json:"locale"`
	InternalSignature string    `json:"internal_signature"`
	CustomerID        string    `json:"customer_id"`
	DeliveryService   string    `json:"delivery_service"`
	Shardkey          string    `json:"shardkey"`
	SmID              int       `json:"sm_id"`
	DateCreated       time.Time `json:"date_created"`
	OofShard          string    `json:"oof_shard"`
}

type Delivery struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Address string `json:"address"`
	Region  string `json:"region"`
	Email   string `json:"email"`
}

func (d *Delivery) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), &d)
}

func (d Delivery) Value() (driver.Value, error) {
	return json.Marshal(d)
}

type Payment struct {
	Transaction  string `json:"transaction"`
	RequestID    string `json:"request_id"`
	Currency     string `json:"currency"`
	Provider     string `json:"provider"`
	Amount       int    `json:"amount"`
	PaymentDT    int    `json:"payment_dt"`
	Bank         string `json:"bank"`
	DeliveryCost int    `json:"delivery_cost"`
	GoodsTotal   int    `json:"goods_total"`
	CustomFee    int    `json:"custom_fee"`
}

func (p *Payment) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), &p)
}

func (p Payment) Value() (driver.Value, error) {
	return json.Marshal(p)
}

type Item struct {
	ChrtID      int    `json:"chrt_id"`
	TrackNumber string `json:"track_number"`
	Price       int    `json:"price"`
	RID         string `json:"rid"`
	Name        string `json:"name"`
	Sale        int    `json:"sale"`
	Size        string `json:"size"`
	TotalPrice  int    `json:"total_price"`
	NmID        int    `json:"nm_id"`
	Brand       string `json:"brand"`
	Status      int    `json:"status"`
}

func (i *Item) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), &i)
}

func (i Item) Value() (driver.Value, error) {
	return json.Marshal(i)
}

type Items []Item

func (i *Items) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), &i)
}

func (i Items) Value() (driver.Value, error) {
	return json.Marshal(i)
}