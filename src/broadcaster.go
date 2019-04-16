package main

import (
	"fmt"
	
	b64 "encoding/base64"
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
	
	if (b.hub.broadcaster == nil) {
		fmt.Println("No broadcaster")
		return 
	}

	sEnc := b64.StdEncoding.EncodeToString(message)

	//fmt.Println(sEnc)

	select {
	case b.hub.broadcast <- []byte(sEnc):
	default:
	}
}
