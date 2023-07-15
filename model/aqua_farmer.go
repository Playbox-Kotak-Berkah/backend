package model

import (
	"github.com/google/uuid"
	"time"
)

type AquaFarmer struct {
	ID         uuid.UUID `gorm:"primaryKey" json:"id"`
	Name       string    `gorm:"not null" json:"name"`
	Phone      string    `gorm:"unique;default:null" json:"phone"` // default is null
	Email      string    `gorm:"unique;not null" json:"email"`
	Password   string    `gorm:"not null" json:"password"`
	IsVerified bool      `gorm:"default:false" json:"is_verified"`
	Picture    string    `gorm:"default:null" json:"picture"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type AquaFarmerRegisterInput struct {
	Name            string `gorm:"binding:required" json:"name"`
	Email           string `gorm:"binding:required" json:"email"`
	Phone           string `gorm:"binding:required" json:"phone"`
	Password        string `gorm:"binding:required" json:"password"`
	ConfirmPassword string `gorm:"binding:required" json:"confirm_password"`
}

type AquaFarmerLoginInput struct {
	Email    string `gorm:"binding:required" json:"email"`
	Password string `gorm:"binding:required" json:"password"`
}

type AquaFarmerEditProfileInput struct {
	Name string `gorm:"binding:required" json:"name"`
}
