package sockets

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  0,
	WriteBufferSize: 0,
}

// SocketUpgradeHandler is main connect ws handler
func SocketUpgradeHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
		return
	}

	startListen(createClient(conn))
}

func startListen(client *SocketClient) {
	defer client.connection.Close()

	for {
		mt, msg, err := client.connection.ReadMessage()
		if err != nil {
			log.Println("ws error read:", err)
			break
		}

		if mt != websocket.TextMessage {
			continue
		}

		log.Printf("ws recv: %s", msg)
		message := &SocketMessage{}
		err = json.Unmarshal(msg, message)

		if err != nil {
			log.Println("ws error parse:", err)
			break
		}

		fire(client, message)
	}
}
