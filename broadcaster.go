package main

import (
	"time"

	"github.com/gorilla/websocket"
)

// Broadcaster sends images for broadcast
type Broadcaster struct {
	GenericClient
}

func (b *Broadcaster) closeConnection() {
	b.hub.broadcaster = nil
	b.conn.Close()
}

func (b *Broadcaster) sendMessage(message []byte) {
	b.hub.broadcast <- message
}

func (b *Broadcaster) writeMessages() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		b.conn.Close()
	}()
	for {
		select {
		case message, ok := <-b.send:
			b.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				b.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := b.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(b.send)
			for i := 0; i < n; i++ {
				w.Write(<-b.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			b.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := b.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}

}
