package handlers

import (
	"net/http"

	"github.com/Virean196/lobby/internal/websocket"
)

func Register(mux *http.ServeMux) {
	mux.HandleFunc("/echo", websocket.Echo)
	mux.Handle("/", http.FileServer(http.Dir("./web")))
}
