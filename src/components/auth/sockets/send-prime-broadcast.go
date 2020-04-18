package sockets

import (
	"audit/src/components/auth/controller"
	"audit/src/sockets"
)

// SendPrimeBroadcast send prime number to all clients
func SendPrimeBroadcast(client sockets.ISocketClient, msg *sockets.SocketMessage) {
	prime, err := controller.GetPrime()
	if err != nil {
		return
	}

	if !msg.ExcludeMe {
		sockets.Broadcast(msg.EventName, prime.String())
	} else {
		predicate := func(clientID string) bool { return clientID != client.GetID() }
		sockets.FilterBroadcast(msg.EventName, prime.String(), predicate)
	}
}
