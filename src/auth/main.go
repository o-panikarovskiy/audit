package auth

import (
	authHttp "audit/src/auth/http"
	authSockets "audit/src/auth/sockets"
	"audit/src/routes/middlewares"
	"audit/src/sockets"
	"net/http"

	"github.com/gorilla/mux"
)

// GetRoutes set auth routes
func GetRoutes(router *mux.Router) {
	sub := router.PathPrefix("/auth").Subrouter()
	sub.HandleFunc("/check", authHttp.CheckSession).Methods("GET")
	sub.Handle("/signup", middlewares.WithJSON(http.HandlerFunc(authHttp.SignUp))).Methods("POST")
	sub.Handle("/signin", middlewares.WithJSON(http.HandlerFunc(authHttp.SignIn))).Methods("POST")
}

// GetSocketEvents set auth routes
func GetSocketEvents() sockets.SocketHandlers {
	return &[]sockets.SocketHandler{
		{Event: "app:prime", Handler: authSockets.SendPrime},
		{Event: "app:prime:broadcast", Handler: authSockets.SendPrimeBroadcast},
	}
}
