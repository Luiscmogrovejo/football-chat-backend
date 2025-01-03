package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	adminpanel "github.com/yourusername/football-chat-backend/adminpanel"
	"github.com/yourusername/football-chat-backend/config"
	"github.com/yourusername/football-chat-backend/controllers"
	"github.com/yourusername/football-chat-backend/middlewares"
)

func main() {
	// 1. Initialize config
	config.InitConfig() // sets up config.Config.DB, etc.

	// 2. Create Gin router
	r := gin.Default()

	// 3. Initialize GoAdmin with our newly created function
	if err := adminpanel.InitializeGoAdmin(r); err != nil {
		log.Fatalf("GoAdmin init failed: %v", err)
	}

	// 4. Example routes
	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)

	r.GET("/ws/:roomID", middlewares.AuthMiddleware(), controllers.UpgradeToWebSocket)

	auth := r.Group("/")
	auth.Use(middlewares.AuthMiddleware())
	{
		auth.GET("/protected", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "You are authenticated"})
		})

		// Rooms
		auth.POST("/rooms", controllers.CreateRoom)
		auth.GET("/rooms", controllers.ListRooms)

		// Matches
		auth.POST("/matches", controllers.CreateMatch)
		auth.GET("/matches", controllers.ListMatches)
		auth.GET("/matches/:id", controllers.GetMatchByID)
		auth.PUT("/matches/:id", controllers.UpdateMatch)
		auth.DELETE("/matches/:id", controllers.DeleteMatch)

		// Predictions
		auth.POST("/predictions", controllers.CreatePrediction)
		auth.GET("/predictions", controllers.ListPredictions)

		// Payment & Subscriptions
		auth.POST("/subscribe", controllers.SubscribeUser)
		auth.POST("/apple-iap-verify", controllers.VerifyAppleIAP)
		auth.POST("/google-iap-verify", controllers.VerifyGoogleIAP)

		// Streaming
		auth.POST("/streams", controllers.CreateStream)
		auth.GET("/streams", controllers.ListStreams)
	}

	// 8. Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server running on :%s\n", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Could not run server: %v", err)
	}
}
