package main

import (
	"log"
	"net/http"

	"github.com/Virean196/lobby/internal/handlers"
	"github.com/Virean196/lobby/internal/hub"
)

func main() {
	mux := http.NewServeMux()

	handlers.Register(mux, hub.NewHub())
	log.Fatal(http.ListenAndServe(":8080", mux))
}
