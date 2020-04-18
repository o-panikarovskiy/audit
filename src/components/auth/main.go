package auth

import (
	"audit/src/components/auth/http"
	authSockets "audit/src/components/auth/sockets"
	"audit/src/sockets"
	"audit/src/utils"
)

// GetRoutes set auth routes
func GetRoutes() utils.RouteHandlers {
	return &[]utils.RouteHandler{
		{Route: "/check", Method: "GET", Handler: http.CheckSession},
	}
}

// GetSocketEvents set auth routes
func GetSocketEvents() sockets.SocketHandlers {
	return &[]sockets.SocketHandler{
		{Event: "app:prime", Handler: authSockets.SendPrime},
		{Event: "app:prime:broadcast", Handler: authSockets.SendPrimeBroadcast},
	}
}
