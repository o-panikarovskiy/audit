package sockets

import (
	"log"

	"github.com/gorilla/websocket"
)

// SocketHandler helps build routes
type SocketHandler struct {
	Event   string
	Handler func(ISocketClient, *SocketMessage)
}

// SocketHandlers is array of RouteHandler
type SocketHandlers *[]SocketHandler

func startListen(client ISocketClient) {
	defer client.Close()

	for {
		mt, bytes, err := client.ReadBytes()
		if err != nil {
			log.Println("ws error read:", err)
			break
		}

		if mt != websocket.TextMessage {
			continue
		}

		log.Printf("ws recv: %s", bytes)
		message, err := client.ToMessage(bytes)

		if err != nil {
			log.Println("ws error parse:", err)
			break
		}

		fireEvents(client, message)
	}
}
