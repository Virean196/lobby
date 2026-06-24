package db

import (
	"database/sql"
	"fmt"
)

type Db struct {
	db *sql.DB
}

func New(db *sql.DB) *Db {
	return &Db{db: db}
}

func (db *Db) CreateUser(username, password string) error {
	_, err := db.db.Exec("INSERT INTO users (username, password) VALUES (?,?)", username, password)
	if err != nil {
		return err
	}
	return nil
}

func (db *Db) GetUser(username string) {
	result := db.db.QueryRow("SELECT username FROM users WHERE username = ?", username)
	fmt.Print(result)
}
