package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	websocketUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			origin := r.Header.Get("origin")
			log.Println("Request from", origin)
			return true
		},
	}
)

type WebsocketManager struct {
	clients      ClientList
	sync.RWMutex // Since many people will concurrently access the API, we need to handle it
}

func NewWebsocketManager() *WebsocketManager {
	return &WebsocketManager{
		clients: make(ClientList),
	}
}

func (m *WebsocketManager) serveWSHelper(w http.ResponseWriter, r *http.Request) {
	log.Println("New Connection")
	authName := r.URL.Query().Get("authName")
	log.Println("Connection from ", authName)
	conn, err := websocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := NewClient(conn, m)

	m.addClient(client)

	go client.readMessages()
	go client.writeMessages()
}

func (m *WebsocketManager) addClient(client *Client) {
	m.Lock()
	defer m.Unlock()
	m.clients[client] = true
}

func (m *WebsocketManager) removeClient(client *Client) {
	m.Lock()
	defer m.Unlock()
	if _, ok := m.clients[client]; ok {
		log.Println("Removing a client")
		client.connection.Close()
		delete(m.clients, client)
	}
}