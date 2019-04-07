package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	//	"github.com/gorilla/websocket"
)

var clientUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 30000,
}

var boatUpgrader = websocket.Upgrader{
	ReadBufferSize:  30000,
	WriteBufferSize: 1024,
}

var clients map[*Client]bool

func serveClient(w http.ResponseWriter, r *http.Request) {
	conn, err := clientUpgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
		return
	}

	client := &Client{conn: conn}

	clients[client] = true

	go client.writeMessages()
	go client.readMessages()

}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./html")))
	http.Handle("/static/", http.FileServer(http.Dir("static/")))

	//	http.Handle("/", server)

	http.HandleFunc("/server", func(w http.ResponseWriter, r *http.Request) {
		serveClient(w, r)
	})

	http.ListenAndServe(":3000", nil)
}
