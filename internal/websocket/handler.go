package websocket

import (
	"log"
	"net/http"

	"github.com/Virean196/lobby/internal/hub"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleWS(h *hub.Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Could not upgrade connection: %v", err)
		return
	}
	username := r.URL.Query().Get("username")
	client := hub.NewClient(conn, username)
	log.Printf("New client connected\nID: %s\nName: %s", client.ID, client.Name)

	go client.ReadLoop(h)
}
