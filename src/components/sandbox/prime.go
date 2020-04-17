package sandbox

import (
	"crypto/rand"
	"fmt"

	"github.com/o-panikarovskiy/audit/src/sockets"
	"github.com/o-panikarovskiy/audit/src/utils"
)

// SendPrime send prime number to client
func SendPrime(client *sockets.SocketClient, msg *sockets.SocketMessage) {
	prime := getPrime()
	client.SendJSON(msg.EventName, utils.HT{"prime": prime, "clentVal": msg.Data})
}

// SendPrimeBroadcast send prime number to all clients
func SendPrimeBroadcast(client *sockets.SocketClient, msg *sockets.SocketMessage) {
	prime := getPrime()
	res := utils.HT{"prime": prime, "clentVal": msg.Data}

	if !msg.ExcludeMe {
		sockets.Broadcast(msg.EventName, res)
	} else {
		predicate := func(clientID string) bool { return clientID != client.ID }
		sockets.FilterBroadcast(msg.EventName, res, predicate)
	}
}

func getPrime() string {
	p, err := rand.Prime(rand.Reader, 100)
	if err != nil {
		return fmt.Sprintf("%v", err)
	}
	return p.String()
}
