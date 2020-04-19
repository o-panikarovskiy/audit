package sockets

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  0,
	WriteBufferSize: 0,
}

// HTTPUpgradeHandler connect socket handler
func HTTPUpgradeHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
		return
	}

	startListen(createClient(conn))
}

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
