package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/football-chat-backend/config"
	"github.com/yourusername/football-chat-backend/models"
)

// Match input struct
type CreateMatchInput struct {
	HomeTeam string `json:"home_team" binding:"required"`
	AwayTeam string `json:"away_team" binding:"required"`
	League   string `json:"league"    binding:"required"`
	DateTime string `json:"date_time" binding:"required"` // store as string for simplicity, convert later
}

// POST /matches
func CreateMatch(c *gin.Context) {
	var input CreateMatchInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	match := models.Match{
		HomeTeam: input.HomeTeam,
		AwayTeam: input.AwayTeam,
		League:   input.League,
		// Parse or store date_time; for simplicity, store as string
		DateTime: input.DateTime,
	}

	if err := config.Config.DB.Create(&match).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, match)
}

// GET /matches
func ListMatches(c *gin.Context) {
	var matches []models.Match
	if err := config.Config.DB.Find(&matches).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, matches)
}

// GET /matches/:id
func GetMatchByID(c *gin.Context) {
	id := c.Param("id")

	var match models.Match
	if err := config.Config.DB.First(&match, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Match not found"})
		return
	}
	c.JSON(http.StatusOK, match)
}

// PUT /matches/:id
func UpdateMatch(c *gin.Context) {
	id := c.Param("id")

	var match models.Match
	if err := config.Config.DB.First(&match, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Match not found"})
		return
	}

	var input CreateMatchInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	match.HomeTeam = input.HomeTeam
	match.AwayTeam = input.AwayTeam
	match.League = input.League
	match.DateTime = input.DateTime

	if err := config.Config.DB.Save(&match).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, match)
}

// DELETE /matches/:id
func DeleteMatch(c *gin.Context) {
	id := c.Param("id")
	var match models.Match

	if err := config.Config.DB.First(&match, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Match not found"})
		return
	}

	if err := config.Config.DB.Delete(&match).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting match"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Match deleted successfully"})
}
