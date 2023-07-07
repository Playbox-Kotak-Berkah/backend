package model

import (
	"github.com/google/uuid"
	"time"
)

type Product struct {
	AquaFarmer   AquaFarmer `gorm:"ForeignKey:AquaFarmerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	AquaFarmerID uuid.UUID  `gorm:"null" json:"aqua_farmer_id"`

	ID          uint    `gorm:"primarykey" json:"id"`
	Name        string  `json:"name"`
	Photo       string  `json:"photo"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Sold        int     `json:"sold"`
	Rating      float64 `json:"rating"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
