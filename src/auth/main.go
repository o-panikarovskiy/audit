package auth

import (
	"audit/src/auth/controller"
	authSockets "audit/src/auth/sockets"
	"audit/src/middlewares"
	"audit/src/sockets"
	"net/http"

	"github.com/gorilla/mux"
)

// GetRoutes set auth routes
func GetRoutes(router *mux.Router) {
	sub := router.PathPrefix("/auth").Subrouter()

	sub.HandleFunc("/check", controller.CheckSession).Methods("GET")

	singUp := middlewares.MdlwJSON(http.HandlerFunc(controller.SignUp))
	signIn := middlewares.MdlwJSON(http.HandlerFunc(controller.SignIn))

	sub.Handle("/signup", singUp).Methods("POST")
	sub.Handle("/signin", signIn).Methods("POST")
}

// GetSocketEvents set auth routes
func GetSocketEvents() sockets.SocketHandlers {
	return &[]sockets.SocketHandler{
		{Event: "app:prime", Handler: authSockets.SendPrime},
		{Event: "app:prime:broadcast", Handler: authSockets.SendPrimeBroadcast},
	}
}
