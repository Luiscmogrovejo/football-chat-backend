package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/football-chat-backend/config"
	"github.com/yourusername/football-chat-backend/models"
)

type CreateRoomInput struct {
	Name string `json:"name" binding:"required"`
}

// POST /rooms
func CreateRoom(c *gin.Context) {
	var input CreateRoomInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Example: get user ID from JWT
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	room := models.Room{
		Name:      input.Name,
		CreatedBy: userID.(uint),
	}

	if err := config.Config.DB.Create(&room).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, room)
}

// GET /rooms
func ListRooms(c *gin.Context) {
	var rooms []models.Room
	if err := config.Config.DB.Find(&rooms).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, rooms)
}
