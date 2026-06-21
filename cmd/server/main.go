package main

import (
	"log"
	"net/http"

	"github.com/Virean196/lobby/internal/handlers"
)

func main() {
	mux := http.NewServeMux()

	handlers.Register(mux)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
