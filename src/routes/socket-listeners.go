package routes

import (
	"github.com/o-panikarovskiy/audit/src/components/sandbox"
	"github.com/o-panikarovskiy/audit/src/sockets"
)

func addSocketEventListeners() {
	sockets.Subscribe("app:prime", sandbox.SendPrime)
	sockets.Subscribe("app:prime:broadcast", sandbox.SendPrimeBroadcast)
}
