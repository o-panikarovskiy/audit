package sockets

// ISocketClient interface
type ISocketClient interface {
	GetID() string
	Close() error
	SendMessage(*SocketMessage) error
	WriteJSON(string, interface{}) error
	ToMessage([]byte) (*SocketMessage, error)
	ReadMessage() (*SocketMessage, error)
	ReadBytes() (int, []byte, error)
}
