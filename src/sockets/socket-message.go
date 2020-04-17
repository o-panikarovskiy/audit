package sockets

// SocketMessage main socket message type
type SocketMessage struct {
	Data      interface{} `json:"data"`
	EventName string      `json:"eventName,omitempty"`
	ClientID  string      `json:"clientId,omitempty"`
	ExcludeMe bool        `json:"excludeMe,omitempty"`
}
