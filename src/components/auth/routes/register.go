package routes

import "github.com/gorilla/mux"

// Register set auth routes
func Register(api *mux.Router) {
	auth := api.PathPrefix("/auth").Subrouter()

	auth.HandleFunc("/check", checkSession)
}
