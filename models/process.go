package models

import (
	"gorm.io/gorm"
)

type Process struct {
	gorm.Model
	Category   string     `json:"category" from:"category"`
	Count      int        `json:"count" from:"count"`
	ArrivesID  int        `json:"arrives_id"`
	Arrives    Arrive     `json:"arrives"`
	Deliveries []Delivery `gorm:"many2many: process_delivery_assoc"`
}
