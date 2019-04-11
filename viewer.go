package main

// Viewer views the broadcast and sends control messages
type Viewer struct {
	GenericClient
}

func (v *Viewer) closeConnection() {
	v.hub.unregister <- v
	v.conn.Close()
}

func (v *Viewer) sendMessage(message []byte) {
	v.hub.broadcaster.send <- message
}
