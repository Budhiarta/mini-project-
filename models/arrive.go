package models

import (
	"gorm.io/gorm"
)

type Arrive struct {
	gorm.Model
	Count     int    `json:"count" from:"count"`
	Suplier   string `json:"suplier" from:"suplier"`
	Price     int    `json:"price" from:"price"`
	Tot_price int    `json:"tot_price" from:"tot_price"`
}
