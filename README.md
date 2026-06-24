# Lobby

A real-time WebSocket lobby server written in Go.

## What it does

- Multiple clients can connect with a username
- Clients can create and join named rooms
- Messages are broadcast to all clients in the same room
- Rooms are automatically created on join and cleaned up when empty
- User registration with hashed passwords

## Tech stack

- Go (net/http, gorilla/websocket, database/sql)
- Vanilla JS + HTML frontend
- Mutex-protected concurrent state management

## Run locally

git clone https://github.com/Virean196/lobby
cd lobby
go run ./cmd/server/main.go

Then open http://localhost:8080 in your browser.

## Project structure

cmd/server    → entrypoint
internal/hub  → Hub, Room, Client structs and logic
internal/handlers → HTTP and WebSocket route registration
internal/db  → MySQL integration for authentication and logging
web           → Frontend HTML/JS client

## Roadmap

- [X] Roster broadcast (show connected clients per room)
- [ ] Disconnect cleanup improvements
- [ ] Room capacity limits
- [ ] Persistent room/player state
- [ ] Authentication System
