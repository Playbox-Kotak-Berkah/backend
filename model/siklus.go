package model

import (
	"github.com/google/uuid"
	"time"
)

type Siklus struct {
	ID            uint    `gorm:"primarykey" json:"id"`
	Tanggal       string  `json:"tanggal"`
	DocSiklus     int     `json:"doc_siklus"`
	PHRealtime    float64 `json:"ph_realtime"`
	PHPagi        float64 `json:"ph_pagi"`
	PHSiang       float64 `json:"ph_siang"`
	PHMalam       float64 `json:"ph_malam"`
	SuhuRealtime  float64 `json:"suhu_realtime"`
	SuhuPagi      float64 `json:"suhu_pagi"`
	SuhuSiang     float64 `json:"suhu_siang"`
	SuhuMalam     float64 `json:"suhu_malam"`
	DORealtime    float64 `json:"do_realtime"`
	DOPagi        float64 `json:"do_pagi"`
	DOSiang       float64 `json:"do_siang"`
	DOMalam       float64 `json:"do_malam"`
	GaramRealtime float64 `json:"garam_realtime"`
	GaramPagi     float64 `json:"garam_pagi"`
	GaramSiang    float64 `json:"garam_siang"`
	GaramMalam    float64 `json:"garam_malam"`

	AquaFarmer   AquaFarmer `gorm:"ForeignKey:AquaFarmerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	AquaFarmerID uuid.UUID  `gorm:"null" json:"user_id"`
	Tambak       Tambak     `gorm:"ForeignKey:TambakID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TambakID     int        `json:"tambak_id"`
	Kolam        Kolam      `gorm:"ForeignKey:KolamID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	KolamID      int        `json:"kolam_id"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
