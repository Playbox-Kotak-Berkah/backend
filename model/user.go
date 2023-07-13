package model

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	Phone     string    `gorm:"unique;default:null" json:"phone"` // default is null
	Email     string    `gorm:"unique;not null" json:"email"`
	Password  string    `gorm:"not null" json:"password"`
	Picture   string    `gorm:"default:null" json:"picture"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserRegisterInput struct {
	Name            string `gorm:"binding:required" json:"name"`
	Email           string `gorm:"binding:required" json:"email"`
	Phone           string `gorm:"binding:required" json:"phone"`
	Password        string `gorm:"binding:required" json:"password"`
	ConfirmPassword string `gorm:"binding:required" json:"confirm_password"`
}

type UserLoginInput struct {
	Email    string `gorm:"binding:required" json:"email"`
	Password string `gorm:"binding:required" json:"password"`
}

type UserUpdateProfileInput struct {
	Name string `json:"name"`
}
