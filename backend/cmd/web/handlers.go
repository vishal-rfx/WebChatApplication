package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home Page"))
}

func serveWS(w http.ResponseWriter, r  *http.Request){
	websocketUpgrader := websocket.Upgrader{
		ReadBufferSize: 1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			origin := r.Header.Get("origin")
			log.Println("Request from", origin)
			return true
		},
	}

	log.Println("Creating a new connection")

	conn, err := websocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := NewClient(conn)

	go client.readMessages()

}