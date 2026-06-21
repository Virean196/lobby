package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	http.HandleFunc("/echo", echo)
	http.Handle("/", http.FileServer(http.Dir("./web")))
	fmt.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func echo(w http.ResponseWriter, r *http.Request) {
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
