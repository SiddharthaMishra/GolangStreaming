package main

import (
	"fmt"
)

// Hub maintains the set of active clients
type Hub struct {

	// Who is broadcasting
	broadcaster *Broadcaster

	// List of open clients
	clients map[*Viewer]bool

	// Register requests from the clients.
	register chan *Viewer

	// Unregister requests from clients.
	unregister chan *Viewer

	// Messages from the ship
	broadcast chan []byte
}

func newHub() *Hub {
	return &Hub{
		register:   make(chan *Viewer, 10),
		unregister: make(chan *Viewer, 10),
		broadcast:  make(chan []byte, 100),
		clients:    make(map[*Viewer]bool),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
			fmt.Println("Connected to the hub!!!")
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message: fmt.Println("sent")
				default:
					client.closeConnection()
				}
			}
		}
	}
}
