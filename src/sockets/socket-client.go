package sockets

import (
	"github.com/gorilla/websocket"
)

// SocketClient app struct
type SocketClient struct {
	ID         string
	connection *websocket.Conn
}

// SendJSON send json data to client
func (client *SocketClient) SendJSON(eventName string, data interface{}) error {
	msg := &SocketMessage{
		Data:      data,
		EventName: eventName,
		ClientID:  client.ID,
	}

	return client.connection.WriteJSON(msg)
}

// Broadcast broadcast event to all clients
func Broadcast(eventName string, data interface{}) error {
	msg := &SocketMessage{
		Data:      data,
		EventName: eventName,
	}

	for _, client := range clients {
		err := client.connection.WriteJSON(msg)
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

	for _, client := range clients {
		if predicate(client.ID) {
			err := client.connection.WriteJSON(msg)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
