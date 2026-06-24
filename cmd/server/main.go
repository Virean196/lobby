package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/Virean196/lobby/internal/db"
	"github.com/Virean196/lobby/internal/handlers"
	"github.com/Virean196/lobby/internal/hub"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	mux := http.NewServeMux()
	dsn := os.Getenv("DBACCESS")
	log.Printf("DSN: %s", dsn)
	dBase, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Print(err)
	}
	database := db.New(dBase)

	h := hub.NewHub(database)

	handlers.Register(mux, h, database)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
