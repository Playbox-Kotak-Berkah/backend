package model

import (
	"github.com/google/uuid"
	"time"
)

type Siklus struct {
	ID            uint   `gorm:"primarykey" json:"id"`
	Tanggal       string `json:"tanggal"`
	DocSiklus     int    `json:"doc_siklus"`
	PHRealtime    int    `json:"ph_realtime"`
	PHPagi        int    `json:"ph_pagi"`
	PHSiang       int    `json:"ph_siang"`
	PHMalam       int    `json:"ph_malam"`
	SuhuRealtime  int    `json:"suhu_realtime"`
	SuhuPagi      int    `json:"suhu_pagi"`
	SuhuSiang     int    `json:"suhu_siang"`
	SuhuMalam     int    `json:"suhu_malam"`
	DORealtime    int    `json:"do_realtime"`
	DOPagi        int    `json:"do_pagi"`
	DOSiang       int    `json:"do_siang"`
	DOMalam       int    `json:"do_malam"`
	GaramRealtime int    `json:"garam_realtime"`
	GaramPagi     int    `json:"garam_pagi"`
	GaramSiang    int    `json:"garam_siang"`
	GaramMalam    int    `json:"garam_malam"`

	AquaFarmer   AquaFarmer `gorm:"ForeignKey:AquaFarmerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	AquaFarmerID uuid.UUID  `gorm:"null" json:"user_id"`
	Tambak       Tambak     `gorm:"ForeignKey:TambakID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TambakID     int        `json:"tambak_id"`
	Kolam        Kolam      `gorm:"ForeignKey:KolamID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	KolamID      int        `json:"kolam_id"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
