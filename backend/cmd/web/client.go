package main

import (
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	connection *websocket.Conn
}

func NewClient(conn *websocket.Conn) *Client {
	return &Client{
		connection: conn,
	}
}

func (c *Client) readMessages() {
	defer func(){
		// cleanup the connection
		c.connection.Close()
	}()
	for { // runs forever
		messageType, payload, err := c.connection.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Unexpected client closed %v", err)
			}
			break
		}
		log.Println(messageType)
		log.Println(string(payload))

	}
}
