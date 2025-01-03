package services

import (
	"log"

	"github.com/yourusername/football-chat-backend/config"
	"github.com/yourusername/football-chat-backend/models"
)

// Example function that triggers notifications for a match event
func NotifyMatchEvent(matchID uint, eventType string) {
	// eventType could be "goal", "yellow_card", etc.

	// Find all users who are "subscribed" to the match. You can define subscription logic as needed.
	// For simplicity, let's assume each user who made a Prediction on this match is "subscribed".
	var predictions []models.Prediction
	if err := config.Config.DB.Where("match_id = ?", matchID).Find(&predictions).Error; err != nil {
		log.Println("Error finding predictions for match:", err)
		return
	}

	// Collect unique user IDs
	userIDs := map[uint]bool{}
	for _, p := range predictions {
		userIDs[p.UserID] = true
	}

	// For each user, get their push token and send a notification
	for userID := range userIDs {
		var user models.User
		if err := config.Config.DB.First(&user, userID).Error; err != nil {
			log.Println("Error finding user:", err)
			continue
		}

		if user.PushToken != "" {
			err := SendExpoPush(user.PushToken, "Match Update", "New event: "+eventType)
			if err != nil {
				log.Println("Error sending push notification:", err)
			}
		}
	}
}
