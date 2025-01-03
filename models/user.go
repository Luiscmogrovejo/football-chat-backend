package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Country   string `gorm:"default:''"`
	FirstName string `gorm:"default:''"`
	LastName  string `gorm:"default:''"`
	Nickname  string `gorm:"uniqueIndex;not null"`
	Email     string `gorm:"uniqueIndex;not null"`
	Password  string `gorm:"not null"`
	AvatarURL string `gorm:"default:''"`
	Bio       string `gorm:"default:''"`
	PushToken string `gorm:"default:''"`
	Plan      string `gorm:"default:'free'"` // e.g. "free", "creator" or "premium"
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
