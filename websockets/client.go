package websockets

import (
	"log"

	"github.com/gorilla/websocket"
	"github.com/yourusername/football-chat-backend/config"
	"github.com/yourusername/football-chat-backend/models"
)

// Client represents a single WebSocket connection.
type Client struct {
	Conn   *websocket.Conn
	SendCh chan []byte
	RoomID string
	UserID uint // user ID from token
}

// ReadPump listens for incoming messages, saves them to DB, and broadcasts them.
func (c *Client) ReadPump(h *Hub) {
	defer func() {
		c.Conn.Close()
	}()

	for {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}

		// Convert msg to string
		content := string(msg)

		// Store in DB
		newMessage := models.Message{
			RoomID:  c.RoomID,
			UserID:  c.UserID,
			Content: content,
		}
		if err := config.Config.DB.Create(&newMessage).Error; err != nil {
			log.Println("Error storing message:", err)
		}

		// Broadcast to other clients in the same room
		for client := range h.Rooms[c.RoomID] {
			if client != c {
				client.SendCh <- msg
			}
		}
	}
}

// WritePump sends messages from the SendCh channel to this WebSocket connection.
func (c *Client) WritePump() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		select {
		case msg, ok := <-c.SendCh:
			if !ok {
				// channel closed
				return
			}
			if err := c.Conn.WriteMessage(websocket.TextMessage, msg); err != nil {
				log.Println("Write error:", err)
				return
			}
		}
	}
}
