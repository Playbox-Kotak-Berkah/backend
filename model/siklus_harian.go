package model

import (
	"time"
)

type SiklusHarian struct {
	ID       uint   `gorm:"primarykey" json:"id"`
	Siklus   Siklus `gorm:"ForeignKey:SiklusID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SiklusID int    `json:"siklus_id"`

	Tanggal       string  `json:"tanggal"`
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

	CreatedAt time.Time
	UpdatedAt time.Time
}

type InputSiklusHarian struct {
	PHPagi     float64 `gorm:"binding:required" json:"ph_pagi"`
	PHSiang    float64 `gorm:"binding:required" json:"ph_siang"`
	PHMalam    float64 `gorm:"binding:required" json:"ph_malam"`
	SuhuPagi   float64 `gorm:"binding:required" json:"suhu_pagi"`
	SuhuSiang  float64 `gorm:"binding:required" json:"suhu_siang"`
	SuhuMalam  float64 `gorm:"binding:required" json:"suhu_malam"`
	DOPagi     float64 `gorm:"binding:required" json:"do_pagi"`
	DOSiang    float64 `gorm:"binding:required" json:"do_siang"`
	DOMalam    float64 `gorm:"binding:required" json:"do_malam"`
	GaramPagi  float64 `gorm:"binding:required" json:"garam_pagi"`
	GaramSiang float64 `gorm:"binding:required" json:"garam_siang"`
	GaramMalam float64 `gorm:"binding:required" json:"garam_malam"`
}
