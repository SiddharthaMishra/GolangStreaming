package main

// Broadcaster sends images for broadcast
type Broadcaster struct {
	GenericClient
}

func (b *Broadcaster) closeConnection() {
	b.hub.broadcaster = nil
	b.conn.Close()
}

func (b *Broadcaster) sendMessage(message []byte) {
	select {
	case b.hub.broadcast <- message:
	default:
	}
}
