package auth

import (
	"audit/src/auth/events"
	"audit/src/auth/handlers"
	"audit/src/middlewares"
	"audit/src/sockets"
	"net/http"

	"github.com/gorilla/mux"
)

// GetRoutes set auth routes
func GetRoutes(router *mux.Router) {
	sub := router.PathPrefix("/auth").Subrouter()

	singUp := middlewares.MdlwRateLimit(middlewares.MdlwJSON(http.HandlerFunc(handlers.SignUp)))
	signIn := middlewares.MdlwRateLimit(middlewares.MdlwJSON(http.HandlerFunc(handlers.SignIn)))
	checkSession := middlewares.MdlwRateLimit(middlewares.MdlwSession(http.HandlerFunc(handlers.CheckSession)))
	signOut := middlewares.MdlwSession(middlewares.MdlwSessionUser(http.HandlerFunc(handlers.SignOut)))
	confirm := middlewares.MdlwRateLimit(http.HandlerFunc(handlers.EndRegistration))

	sub.Handle("/signup", singUp).Methods("POST")
	sub.Handle("/signin", signIn).Methods("POST")
	sub.Handle("/signout", signOut).Methods("POST")
	sub.Handle("/check", checkSession).Methods("GET")
	sub.Handle("/confirm/{token}", confirm).Methods("GET")
}

// GetSocketEvents set auth routes
func GetSocketEvents() sockets.SocketHandlers {
	return &[]sockets.SocketHandler{
		{Event: "app:prime", Handler: events.SendPrime},
		{Event: "app:prime:broadcast", Handler: events.SendPrimeBroadcast},
	}
}
