package main

import (
	"sync"

	"github.com/gorilla/websocket"
)

// Client handles a connection
type Client struct {
	conn *websocket.Conn

	// locks when writing to prevent parallel writes
	mutex sync.Mutex
}

func (*Client) readMessages() {

}

func (*Client) writeMessages() {

}
