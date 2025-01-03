package models

import (
	"time"

	"gorm.io/gorm"
)

type Stream struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string `gorm:"not null"`
	MatchID   uint   `gorm:"index;not null"`
	StreamURL string `gorm:"not null"`
	CreatedBy uint   `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
