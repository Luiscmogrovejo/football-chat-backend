package models

import (
	"time"

	"gorm.io/gorm"
)

// Prediction is a guess the user makes during a match
type Prediction struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"index;not null"`
	MatchID   uint   `gorm:"index;not null"`
	PredType  string `gorm:"not null"`  // e.g. "next_goal_scorer", "next_card"
	PredData  string `gorm:"not null"`  // e.g. "PlayerName" or "TeamName"
	Points    int    `gorm:"default:0"` // Points earned
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
