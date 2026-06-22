# Lobby

A real-time WebSocket lobby server written in Go.

## What it does

- Multiple clients can connect with a username
- Clients can create and join named rooms
- Messages are broadcast to all clients in the same room
- Rooms are automatically created on join and cleaned up when empty

## Tech stack

- Go (net/http, gorilla/websocket)
- Vanilla JS + HTML frontend
- Mutex-protected concurrent state management

## Run locally

git clone https://github.com/Virean196/lobby
cd lobby
go run ./cmd/server

Then open http://localhost:8080 in your browser.

## Project structure

cmd/server    → entrypoint
internal/hub  → Hub, Room, Client structs and logic
internal/handlers → HTTP and WebSocket route registration
web           → Frontend HTML/JS client

## Roadmap

- [ ] Roster broadcast (show connected clients per room)
- [ ] Disconnect cleanup improvements
- [ ] Room capacity limits
- [ ] Persistent room/player state
