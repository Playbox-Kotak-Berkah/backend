package model

import (
	"github.com/google/uuid"
	"time"
)

type Kolam struct {
	ID                uint   `gorm:"primarykey" json:"id"`
	Name              string `json:"name"`
	LampuTambakStatus bool   `json:"lampu_tambak_status"`
	KincirAirStatus   bool   `json:"kincir_air_status"`
	KeranAirStatus    bool   `json:"keran_air_status"`

	AquaFarmer   AquaFarmer `gorm:"ForeignKey:AquaFarmerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	AquaFarmerID uuid.UUID  `gorm:"null" json:"aqua_farmer_id"`
	Tambak       Tambak     `gorm:"ForeignKey:TambakID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TambakID     int        `json:"tambak_id"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
