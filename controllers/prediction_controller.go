package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/football-chat-backend/config"
	"github.com/yourusername/football-chat-backend/models"
)

// Example input struct
type CreatePredictionInput struct {
	MatchID  uint   `json:"match_id" binding:"required"`
	PredType string `json:"pred_type" binding:"required"` // e.g., "next_goal_scorer", "next_card"
	PredData string `json:"pred_data"`                    // e.g., "PlayerName" or "TeamName"
}

// POST /predictions
func CreatePrediction(c *gin.Context) {
	var input CreatePredictionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	prediction := models.Prediction{
		UserID:   userID.(uint),
		MatchID:  input.MatchID,
		PredType: input.PredType,
		PredData: input.PredData,
	}

	if err := config.Config.DB.Create(&prediction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, prediction)
}

// GET /predictions
func ListPredictions(c *gin.Context) {
	var predictions []models.Prediction
	if err := config.Config.DB.Find(&predictions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, predictions)
}
