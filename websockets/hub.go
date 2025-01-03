package websockets

// Hub manages multiple rooms
type Hub struct {
	Rooms map[string]map[*Client]bool
}

// global instance
var H = NewHub()

func NewHub() *Hub {
	return &Hub{
		Rooms: make(map[string]map[*Client]bool),
	}
}

// JoinRoom adds a client to a room
func (h *Hub) JoinRoom(roomID string, client *Client) {
	if h.Rooms[roomID] == nil {
		h.Rooms[roomID] = make(map[*Client]bool)
	}
	h.Rooms[roomID][client] = true
}
