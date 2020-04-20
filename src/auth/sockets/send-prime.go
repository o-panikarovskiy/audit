package sockets

import (
	"audit/src/sockets"
)

// SendPrime send prime number to client
func SendPrime(client sockets.ISocketClient, msg *sockets.SocketMessage) {
	client.WriteJSON(msg.EventName, "13")
}
