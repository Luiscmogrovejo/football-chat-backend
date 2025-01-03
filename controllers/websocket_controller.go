package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/yourusername/football-chat-backend/websockets"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // In production, refine for security
	},
}

// UpgradeToWebSocket upgrades an HTTP request to a WebSocket connection
// GET /ws/:roomID
func UpgradeToWebSocket(c *gin.Context) {
	// Retrieve user_id from context
	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized - missing user_id"})
		return
	}
	roomID := c.Param("roomID")

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upgrade WebSocket"})
		return
	}

	client := &websockets.Client{
		Conn:   conn,
		SendCh: make(chan []byte),
		RoomID: roomID,
		UserID: userID.(uint),
	}

	// Join the global Hub
	websockets.H.JoinRoom(roomID, client)

	go client.WritePump()
	go client.ReadPump(websockets.H)
}
