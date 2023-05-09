package models

import (
	"gorm.io/gorm"
)

type Delivery struct {
	gorm.Model
	Client    string    `json:"client" from:"client"`
	Count     int       `json:"count" from:"count"`
	Price     int       `json:"price" from:"price"`
	Tot_price int       `json:"tot_price" from:"tot_price"`
	Process   []Process `gorm:"many2many: process_delivery_assoc"`
}
