package models

import (
	"time"

	"gorm.io/gorm"
)

// Match model to store match data
type Match struct {
	ID        uint   `gorm:"primaryKey"`
	HomeTeam  string `gorm:"not null"`
	AwayTeam  string `gorm:"not null"`
	League    string `gorm:"not null"`
	DateTime  string `gorm:"not null"` // For simplicity, using string. Could be time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
