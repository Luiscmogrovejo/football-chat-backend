package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/football-chat-backend/config"
	"github.com/yourusername/football-chat-backend/models"
)

type CreateStreamInput struct {
	Title     string `json:"title" binding:"required"`
	MatchID   uint   `json:"match_id" binding:"required"`
	StreamURL string `json:"stream_url" binding:"required"`
}

// POST /streams
func CreateStream(c *gin.Context) {
	userID := c.GetUint("user_id")

	var input CreateStreamInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	stream := models.Stream{
		Title:     input.Title,
		MatchID:   input.MatchID,
		StreamURL: input.StreamURL,
		CreatedBy: userID,
	}

	if err := config.Config.DB.Create(&stream).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, stream)
}

// GET /streams
func ListStreams(c *gin.Context) {
	var streams []models.Stream
	if err := config.Config.DB.Find(&streams).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, streams)
}
