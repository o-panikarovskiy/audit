package sockets

// SocketHandler helps build routes
type SocketHandler struct {
	Event   string
	Handler func(ISocketClient, *SocketMessage)
}

// SocketHandlers is array of RouteHandler
type SocketHandlers *[]SocketHandler
