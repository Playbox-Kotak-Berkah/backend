package model

type Paket struct {
	ID    uint   `gorm:"primarykey" json:"id"`
	Name  string `json:"name"`
	Price int64  `json:"price"`
	Month int    `json:"month"`
}
