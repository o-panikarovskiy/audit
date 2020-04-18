package core

import (
	"audit/src/components/auth"
	"audit/src/config"
	"audit/src/sockets"
)

func addSocketEventListeners(cfg *config.AppConfig) {
	sockets.SubscribeEvents(auth.GetSocketEvents())
}
