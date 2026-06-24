package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Virean196/lobby/internal/db"
	"github.com/Virean196/lobby/internal/hub"
	"github.com/Virean196/lobby/internal/websocket"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(mux *http.ServeMux, h *hub.Hub, db *db.Db) {
	mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		websocket.HandleWS(h, w, r)
	})
	mux.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		var request RegisterRequest
		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"message": "invalid request"}`))
		}
		err = json.Unmarshal(body, &request)
		if err != nil {
		}
		err = db.CreateUser(request.Username, request.Password)
		fmt.Printf("User: %s - Created", request.Username)
		if err != nil {
			log.Printf("Could not create user: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write([]byte(`{"message": "user created"}`))
	})
	mux.Handle("/", http.FileServer(http.Dir("./web")))
}
