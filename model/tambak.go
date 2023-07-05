package model

import (
	"github.com/google/uuid"
	"time"
)

type Tambak struct {
	ID           uint       `gorm:"primaryKey" json:"id"`
	Name         string     `json:"name"`
	AquaFarmer   AquaFarmer `gorm:"ForeignKey:AquaFarmerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	AquaFarmerID uuid.UUID  `gorm:"null" json:"aqua_farmer_id"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
