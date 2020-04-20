package events

import (
	"audit/src/sockets"
)

// SendPrimeBroadcast send prime number to all clients
func SendPrimeBroadcast(client sockets.ISocketClient, msg *sockets.SocketMessage) {
	prime := "123"

	if !msg.ExcludeMe {
		sockets.Broadcast(msg.EventName, prime)
	} else {
		predicate := func(clientID string) bool { return clientID != client.GetID() }
		sockets.FilterBroadcast(msg.EventName, prime, predicate)
	}
}
