package handlers

import (
	"net/http"

	"github.com/Virean196/lobby/internal/db"
	"github.com/Virean196/lobby/internal/hub"
	"github.com/Virean196/lobby/internal/websocket"
)

func Register(mux *http.ServeMux, h *hub.Hub, db *db.Db) {
	mux.HandleFunc("/echo", websocket.Echo)
	mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		websocket.HandleWS(h, w, r)
	})
	mux.Handle("/", http.FileServer(http.Dir("./web")))
}
