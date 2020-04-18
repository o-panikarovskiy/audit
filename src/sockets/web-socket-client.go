package sockets

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

type socketClient struct {
	ID         string
	connection *websocket.Conn
}

// Close send json data to client
func (client *socketClient) GetID() string {
	return client.ID
}

// Close send json data to client
func (client *socketClient) Close() error {
	return client.connection.Close()
}

// SendMessage writes the JSON encoding of v as a message.
//
// See the documentation for encoding/json Marshal for details about the
// conversion of Go values to JSON.
func (client *socketClient) SendMessage(message *SocketMessage) error {
	return client.connection.WriteJSON(message)
}

// WriteJSON writes the JSON encoding of v as a message.
//
// See the documentation for encoding/json Marshal for details about the
// conversion of Go values to JSON.
func (client *socketClient) WriteJSON(eventName string, data interface{}) error {
	msg := &SocketMessage{
		Data:      data,
		EventName: eventName,
		ClientID:  client.ID,
	}

	return client.connection.WriteJSON(msg)
}

// ReadMessage returns *SocketMessage
func (client *socketClient) ReadMessage() (*SocketMessage, error) {
	mt, bytes, err := client.ReadBytes()
	if err != nil {
		return nil, err
	}

	if mt != websocket.TextMessage {
		return nil, fmt.Errorf("message type is not text")
	}

	return client.ToMessage(bytes)
}

// ToMessage converts butes to *SocketMessage
func (client *socketClient) ToMessage(bytes []byte) (*SocketMessage, error) {
	message := &SocketMessage{}
	err := json.Unmarshal(bytes, message)

	if err != nil {
		return nil, err
	}

	return message, nil
}

// ReadBytes is a helper method for getting a reader using NextReader and
// reading from that reader to a buffer.
func (client *socketClient) ReadBytes() (int, []byte, error) {
	return client.connection.ReadMessage()
}
