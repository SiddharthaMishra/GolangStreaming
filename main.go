package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

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

func serveBroadcaster(w http.ResponseWriter, r *http.Request, h *Hub) *Hub {
	conn, err := clientUpgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
		return nil
	}

	b := &Broadcaster{GenericClient: makeWS(conn, h)}

	log.Printf("in func broadcaster: %+v\n", b)
	h.broadcaster = b

	log.Printf("in func : %+v\n", h)

	return h
}

func serveViewer(w http.ResponseWriter, r *http.Request, h *Hub) *Viewer {
	conn, err := clientUpgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
		return nil
	}

	v := &Viewer{GenericClient: makeWS(conn, h)}
	v.hub.register <- v

	return v
}

func main() {
	f, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

	if err != nil {
		fmt.Println("hii")
		log.Fatal(err)
	}

	defer f.Close()
	log.SetOutput(f)

	log.Println("")

	http.Handle("/", http.FileServer(http.Dir("./html")))

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	h := newHub()

	http.HandleFunc("/viewer", func(w http.ResponseWriter, r *http.Request) {
		if h != nil {
			v := serveViewer(w, r, h)

			go readMessages(v)
			go writeMessages(v)
		}
	})

	http.HandleFunc("/broadcaster", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("boat pinged")
		if h.broadcaster == nil {
			serveBroadcaster(w, r, h)

			log.Printf("%+v\n", h)
			log.Printf("Broadcaster: %+v\n", h.broadcaster)

			go h.run()
			go readMessages(h.broadcaster)
			go writeMessages(h.broadcaster)
		}
		fmt.Printf("%+v\n", h)
	})

	http.ListenAndServe(":3000", nil)
}
