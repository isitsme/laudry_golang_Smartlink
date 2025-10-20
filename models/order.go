package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	CustomerName string    `gorm:"size:100;not null" json:"customer_name"`
	CustomerID   uint      `json:"customer_id"`
	ServiceID    uint      `json:"service_id"`
	Service      Service   `gorm:"foreignKey:ServiceID" json:"service"`
	Quantity     int       `gorm:"not null" json:"quantity"`
	TotalPrice   float64   `gorm:"not null" json:"total_price"`
	Status       string    `gorm:"size:50;default:'pending'" json:"status"` // pending, processing, done, cancelled
	DueDate      time.Time `json:"due_date"`
}
