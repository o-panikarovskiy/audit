package sockets

import (
	"audit/src/utils"
	"fmt"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

var connections sync.Map

// GetClient return client connection by id
func GetClient(clientID string) (ISocketClient, error) {
	val, ok := connections.Load(clientID)
	if !ok {
		return nil, fmt.Errorf("Socket client not found. Client ID: %v", clientID)
	}

	return val.(ISocketClient), nil
}

// Broadcast broadcast event to all clients
func Broadcast(eventName string, data interface{}) {
	msg := &SocketMessage{
		Data:      data,
		EventName: eventName,
	}

	connections.Range(func(key interface{}, val interface{}) bool {
		(val.(ISocketClient)).SendMessage(msg)
		return true
	})
}

// FilterBroadcast broadcast event to clients by predicate
func FilterBroadcast(eventName string, data interface{}, predicate func(clientID string, userID string) bool) {
	msg := &SocketMessage{
		Data:      data,
		EventName: eventName,
	}

	connections.Range(func(key interface{}, val interface{}) bool {
		client := (val.(ISocketClient))
		if predicate(client.GetID(), client.GetUserID()) {
			client.SendMessage(msg)
		}
		return true
	})
}

// FindClient broadcast event to clients by predicate
func FindClient(predicate func(clientID string, userID string) bool) ISocketClient {
	var res ISocketClient

	connections.Range(func(key interface{}, val interface{}) bool {
		client := (val.(ISocketClient))
		if predicate(client.GetID(), client.GetUserID()) {
			res = client
			return false
		}
		return true
	})

	return res
}

// RemoveClient func
func RemoveClient(clientID string) {
	connections.Delete(clientID)
}

func createClient(conn *websocket.Conn, userID string) ISocketClient {
	client := &socketClient{
		connection: conn,
		ID:         utils.CreateGUID(),
		UserID:     userID,
	}

	conn.SetCloseHandler(func(code int, text string) error {
		log.Println("Client disconnected", client.GetID())
		RemoveClient(client.GetID())
		return nil
	})

	log.Println("Client connected", client.GetID())
	connections.Store(client.GetID(), client)

	client.WriteJSON("socket:client:id", client.GetID())

	return client
}
