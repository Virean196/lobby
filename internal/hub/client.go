package hub

import (
	"encoding/json"
	"log"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	Conn *websocket.Conn
	ID   string
	Name string
	Room string
}

func NewClient(conn *websocket.Conn, username string) *Client {
	return &Client{
		ID:   uuid.New().String(),
		Conn: conn,
		Name: username,
		Room: "",
	}
}

func (c *Client) Send(msg Message) error {
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	return c.Conn.WriteMessage(websocket.TextMessage, data)
}

func (c *Client) ReadLoop(h *Hub) {
	defer func() {
		h.LeaveRoom(c, c.Room)
	}()
	for {
		_, raw, err := c.Conn.ReadMessage()
		if err != nil {
			log.Printf("Client %s disconnected", c.ID)
			break
		}
		var msg Message
		err = json.Unmarshal(raw, &msg)
		if err != nil {
			log.Printf("Could not parse message: %v", err)
			continue
		}
		switch msg.Type {
		case "join":
			h.JoinRoom(c, msg.Room)
		case "leave":
			h.LeaveRoom(c, msg.Room)
		case "message":
			h.Broadcast(h.Rooms[msg.Room], msg.Content)
		}
	}
}
