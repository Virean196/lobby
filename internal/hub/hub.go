package hub

import (
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/Virean196/lobby/internal/db"
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
	Db    *db.Db
}

func NewHub(db *db.Db) *Hub {
	return &Hub{
		Rooms: map[string]*Room{},
		Db:    db,
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
	hub.listUsersInRoom(room)
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
	hub.listUsersInRoom(room)
	return nil
}

func (hub *Hub) Broadcast(room *Room, message string, sender string) {
	hub.mu.Lock()
	defer hub.mu.Unlock()
	for _, client := range room.Clients {
		msg := Message{
			Type:    "message",
			Room:    room.Name,
			Content: message,
			Sender:  sender,
		}
		err := client.Send(msg)
		if err != nil {
			log.Printf("Could not send message to client %s: %v", client.Name, err)
		}
	}
}

func (hub *Hub) listUsersInRoom(room *Room) {
	var userList []string
	for _, client := range room.Clients {
		userList = append(userList, client.Name)
	}
	usersCSV := strings.Join(userList, ",")
	for _, client := range room.Clients {
		err := client.Send(Message{Type: "roster", Content: usersCSV})
		if err != nil {
			log.Printf("Could not send userlist message: %v", err)
		}
	}
}
