package sockets

import (
	"fmt"
	"log"

	"audit/src/utils"

	"github.com/gorilla/websocket"
)

var clients map[string]*SocketClient = make(map[string]*SocketClient)

// GetClient return client connection by id
func GetClient(clientID string) (*SocketClient, error) {
	conn, ok := clients[clientID]
	if !ok {
		return nil, fmt.Errorf("Socket client not found. Client ID: %v", clientID)
	}

	return conn, nil
}

func createClient(conn *websocket.Conn) *SocketClient {
	clientID := utils.CreateGUID()

	conn.SetCloseHandler(func(code int, text string) error {
		log.Println("Client disconnected", clientID)
		removeClient(clientID)
		return nil
	})

	client := &SocketClient{ID: clientID, connection: conn}
	clients[clientID] = client

	log.Println("Client connected", clientID)
	client.SendJSON("socket:client:id", clientID)

	return client
}

func removeClient(clientID string) {
	delete(clients, clientID)
}
