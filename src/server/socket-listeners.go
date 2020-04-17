package server

import (
	"audit/src/components/sandbox"
	"audit/src/sockets"
)

func addSocketEventListeners() {
	sockets.Subscribe("app:prime", sandbox.SendPrime)
	sockets.Subscribe("app:prime:broadcast", sandbox.SendPrimeBroadcast)
}
