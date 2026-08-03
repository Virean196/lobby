package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/Virean196/lobby/internal/db"
	"github.com/Virean196/lobby/internal/hub"
	"github.com/Virean196/lobby/internal/websocket"
)

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(mux *http.ServeMux, h *hub.Hub, db *db.Db) {
	mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		websocket.HandleWS(h, w, r)
	})
	mux.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		var request UserRequest
		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"message": "invalid request"}`))
		}
		err = json.Unmarshal(body, &request)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"message": "unable to unmarshal request body"}`))
		}
		err = db.CreateUser(request.Username, request.Password)
		loweredUsername := strings.ToLower(request.Username)
		fmt.Printf("User: %s - Created\n", loweredUsername)
		if err != nil {
			log.Printf("Could not create user: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "invalid username"`))
			return
		}
		w.Write([]byte(`{"message": "user created"}`))
	})

	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		var request UserRequest
		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"message": "invalid request"}`))
		}
		err = json.Unmarshal(body, &request)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"message": "unable to unmarshal request body"}`))
		}
		loweredUsername := strings.ToLower(request.Username)
		err = db.Login(loweredUsername, request.Password)
		if err == nil {
			w.WriteHeader(200)
			w.Write([]byte(`{"message": "logged in successfully!"}`))
			fmt.Printf("User: %s - Logged in successfully!\n", loweredUsername)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"message": "invalid login request"}`))
		}
	})

	// Host a fileserver on the folder "web"
	mux.Handle("/", http.FileServer(http.Dir("./web")))
}
