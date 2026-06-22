package hub

import (
	"fmt"
	"log"
	"sync"

	"github.com/google/uuid"
)

type Message struct {
	Type    string `json:"type"`
	Room    string `json:"room"`
	Content string `json:"content"`
	Sender  string `json:"sender"`
}

type Room struct {
	ID      string
	Name    string
	Clients map[string]*Client
}

type Hub struct {
	Rooms map[string]*Room
	mu    sync.Mutex
}

func NewHub() *Hub {
	return &Hub{
		Rooms: map[string]*Room{},
	}
}

func (hub *Hub) JoinRoom(client *Client, roomName string) Message {
	hub.mu.Lock()
	defer hub.mu.Unlock()
	if client.Room != "" {
		hub.leaveRoom(client, client.Room)
	}
	room, ok := hub.Rooms[roomName]
	if !ok {
		room = &Room{
			ID:      uuid.New().String(),
			Name:    roomName,
			Clients: map[string]*Client{},
		}
		hub.Rooms[roomName] = room
		log.Printf("Room Created!\nID: %s\nName: %s\n", room.ID, room.Name)
	}
	room.Clients[client.ID] = client
	client.Room = roomName
	log.Printf("Client: %s\nAdded to Room: %s", client.Name, room.Name)
	return Message{
		Type:    "joined",
		Room:    roomName,
		Content: "",
	}
}

func (h *Hub) LeaveRoom(client *Client, roomName string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.leaveRoom(client, roomName)
}

func (hub *Hub) leaveRoom(client *Client, roomName string) error {
	room, ok := hub.Rooms[roomName]
	if !ok {
		return fmt.Errorf("No such room")
	}
	delete(room.Clients, client.ID)
	log.Printf("Client: %s\nRemoved from room: %s", client.Name, room.Name)
	if len(room.Clients) == 0 {
		delete(hub.Rooms, roomName)
		log.Printf("Room: %s is empty, deleting!", room.Name)
	}
	return nil
}

func (hub *Hub) Broadcast(room *Room, message string, sender string) {
	hub.mu.Lock()
	defer hub.mu.Unlock()
	for client := range room.Clients {
		msg := Message{
			Type:    "message",
			Room:    room.Name,
			Content: message,
			Sender:  sender,
		}
		err := room.Clients[client].Send(msg)
		if err != nil {
			log.Printf("Could not send message to client %s: %v", client, err)
		}
	}
}
