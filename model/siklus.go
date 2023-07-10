package model

import (
	"github.com/google/uuid"
	"time"
)

type Siklus struct {
	ID           uint       `gorm:"primarykey" json:"id"`
	AquaFarmer   AquaFarmer `gorm:"ForeignKey:AquaFarmerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	AquaFarmerID uuid.UUID  `gorm:"null" json:"aqua_farmer_id"`
	Tambak       Tambak     `gorm:"ForeignKey:TambakID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TambakID     int        `json:"tambak_id"`
	Kolam        Kolam      `gorm:"ForeignKey:KolamID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	KolamID      int        `json:"kolam_id"`

	Name      string `json:"name"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

type SiklusInput struct {
	Name string `gorm:"binding:required" json:"name"`
}
