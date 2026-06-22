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

func Echo(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrader error: ", err)
		return
	}
	defer conn.Close()
	for {
		msgType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read error: ", err)
			break
		}
		log.Printf("received: %s", message)
		err = conn.WriteMessage(msgType, message)
		if err != nil {
			log.Println("write error: ", err)
			break
		}
	}
}

func HandleWS(h *hub.Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Could not upgrade connection: %v", err)
		return
	}
	username := r.URL.Query().Get("username")
	var client *hub.Client
	if len(username) >= 3 {
		client = hub.NewClient(conn, username)
	}
	log.Printf("New client connected\nID: %s\nName: %s", client.ID, client.Name)

	go client.ReadLoop(h)
}
