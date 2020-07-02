package routes

import (
	"net/http"

	"audit/src/auth"
	"audit/src/config"
	"audit/src/middlewares"
	"audit/src/sockets"

	"github.com/gorilla/mux"
)

// CreateRouter create main http.Handler
func CreateRouter(cfg *config.AppConfig) http.Handler {
	router := mux.NewRouter()

	socketHandler := middlewares.MdlwSession(middlewares.MdlwSessionUser(http.HandlerFunc(sockets.HTTPUpgradeHandler)))

	api := router.PathPrefix("/api").Subrouter()
	api.Use(middlewares.MdlwError)
	api.Use(middlewares.MdlwLog)
	api.Use(middlewares.MdlwTypedContext)

	api.Handle("/ws", socketHandler).Methods("GET")
	api.HandleFunc("/health", health).Methods("GET")

	auth.GetRoutes(api)

	api.NotFoundHandler = http.HandlerFunc(notFound)
	api.MethodNotAllowedHandler = http.HandlerFunc(notAllowed)

	router.PathPrefix("/").HandlerFunc(createSpaHandler(cfg.StaticDir, "index.html"))
	return router
}
