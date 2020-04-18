package sockets

import (
	"audit/src/components/auth/controller"
	"audit/src/sockets"
)

// SendPrime send prime number to client
func SendPrime(client sockets.ISocketClient, msg *sockets.SocketMessage) {
	prime, err := controller.GetPrime()
	if err != nil {
		return
	}
	client.WriteJSON(msg.EventName, prime.String())
}
