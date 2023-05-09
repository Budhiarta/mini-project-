package dto

import (
	"miniProject/models"

	"gorm.io/gorm"
)

type CreateDeliveryRequest struct {
	Client   string           `json:"client" `
	Price    int              `json:"price"`
	Proceses []ProcessRequest `json:"proceses"`
}

type ProcessRequest struct {
	Id    uint `json:"id"`
	Count int  `json:"count"`
}

func (c *CreateDeliveryRequest) ToEntity() *models.Delivery {
	var process []models.Process
	var count int

	for _, each := range c.Proceses {
		process = append(process, *each.ToEntity())
		count += each.Count

	}

	return &models.Delivery{
		Client:    c.Client,
		Price:     c.Price,
		Count:     count,
		Tot_price: count * c.Price,
		Process:   process,
	}
}

func (p *ProcessRequest) ToEntity() *models.Process {
	return &models.Process{
		Model: gorm.Model{ID: p.Id},
	}
}
