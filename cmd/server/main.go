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
)

func main() {
	mux := http.NewServeMux()

	dbConnString := os.Getenv("DBACCESS")
	dBase, err := sql.Open("mysql", dbConnString)
	if err != nil {
		log.Print(err)
	}
	database := db.New(dBase)

	h := hub.NewHub(database)

	handlers.Register(mux, h, database)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
