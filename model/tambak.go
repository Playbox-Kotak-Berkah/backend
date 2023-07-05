package model

import (
	"github.com/google/uuid"
	"time"
)

type Tambak struct {
	ID           uint       `gorm:"primaryKey" json:"id"`
	Name         string     `json:"name"`
	AquaFarmer   AquaFarmer `gorm:"foreignKey:AquaFarmerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"aqua_farmer"`
	AquaFarmerID uuid.UUID  `gorm:"type:uuid;index;null" json:"aqua_farmer_id"`
	CreatedAt    time.Time  `json:"-"`
	UpdatedAt    time.Time  `json:"-"`
}

type AddTambak struct {
	Name string `gorm:"binding:required" json:"name"`
}
