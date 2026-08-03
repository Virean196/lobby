package db

import (
	"database/sql"
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type Db struct {
	db *sql.DB
}

func New(db *sql.DB) *Db {
	return &Db{db: db}
}

func (db *Db) CreateUser(username, password string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	loweredUserName := strings.ToLower(username)
	if err != nil {
		return err
	}
	_, err = db.db.Exec("INSERT INTO users (username, password) VALUES (?,?)", loweredUserName, hashed)
	if err != nil {
		return err
	}
	return nil
}

func (db *Db) getUserPassword(username string) (string, error) {
	var password string
	result := db.db.QueryRow("SELECT password FROM users WHERE username = ?", username)
	err := result.Scan(&password)
	if err != nil {
		return "", err
	}
	return password, nil
}

func (db *Db) Login(username, password string) error {
	var pw string
	loweredUserName := strings.ToLower(username)
	pw, err := db.getUserPassword(loweredUserName)
	if err != nil {
		fmt.Printf("No rows: %v", err)
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(pw), []byte(password))
	if err != nil {
		fmt.Printf("Compare pw: %v", err)
		return err
	}
	return nil
}
