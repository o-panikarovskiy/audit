package server

import (
	"audit/src/auth"
	"audit/src/config"
	"audit/src/sockets"
)

func addSocketEventListeners(cfg *config.AppConfig) {
	sockets.SubscribeEvents(auth.GetSocketEvents())
}
