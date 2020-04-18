package sockets

import (
	"audit/src/utils"
	"fmt"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type connectionsMap struct {
	sync.Mutex
	clients map[string]ISocketClient
}

var connections = &connectionsMap{clients: make(map[string]ISocketClient)}

// GetClient return client connection by id
func GetClient(clientID string) (ISocketClient, error) {
	conn, ok := connections.clients[clientID]
	if !ok {
		return nil, fmt.Errorf("Socket client not found. Client ID: %v", clientID)
	}

	return conn, nil
}

// Broadcast broadcast event to all clients
func Broadcast(eventName string, data interface{}) error {
	msg := &SocketMessage{
		Data:      data,
		EventName: eventName,
	}

	for _, client := range connections.clients {
		err := client.SendMessage(msg)
		if err != nil {
			return err
		}
	}

	return nil
}

// FilterBroadcast broadcast event to clients by predicate
func FilterBroadcast(eventName string, data interface{}, predicate func(clientID string) bool) error {
	msg := &SocketMessage{
		Data:      data,
		EventName: eventName,
	}

	for _, client := range connections.clients {
		if predicate(client.GetID()) {
			err := client.SendMessage(msg)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func createClient(conn *websocket.Conn) ISocketClient {
	client := &socketClient{
		connection: conn,
		ID:         utils.CreateGUID(),
	}

	conn.SetCloseHandler(func(code int, text string) error {
		log.Println("Client disconnected", client.GetID())
		removeClient(client.GetID())
		return nil
	})

	connections.Lock()
	defer connections.Unlock()
	connections.clients[client.GetID()] = client

	log.Println("Client connected", client.GetID())
	client.WriteJSON("socket:client:id", client.GetID())

	return client
}

func removeClient(clientID string) {
	connections.Lock()
	defer connections.Unlock()
	delete(connections.clients, clientID)
}
