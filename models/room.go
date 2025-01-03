package models

import (
	"time"

	"gorm.io/gorm"
)

// Room represents a chat room
type Room struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"uniqueIndex;not null"`
	CreatedBy uint   `gorm:"not null"` // ID of the user who created the room
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
