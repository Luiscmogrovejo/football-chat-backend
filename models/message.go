package models

import (
	"time"

	"gorm.io/gorm"
)

// Message stores individual chat messages
type Message struct {
	ID        uint   `gorm:"primaryKey"`
	RoomID    string `gorm:"index;not null"`
	UserID    uint   `gorm:"index;not null"`
	Content   string `gorm:"type:text;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
