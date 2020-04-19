package routes

import (
	"net/http"

	"audit/src/auth"
	"audit/src/config"
	"audit/src/context"
	"audit/src/sockets"

	"github.com/gorilla/mux"
)

// CreateRouter create main http.Handler
func CreateRouter(cfg *config.AppConfig) http.Handler {
	router := mux.NewRouter()

	api := router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/health", health).Methods("GET")
	api.HandleFunc("/ws", sockets.HTTPUpgradeHandler).Methods("GET")

	api.Use(WithError)
	api.Use(context.WithTypedContext)
	api.Use(WithJSON)
	api.Use(WithLog)

	setSubRoutes(api, "/auth", auth.GetRoutes())

	api.NotFoundHandler = http.HandlerFunc(notFound)
	api.MethodNotAllowedHandler = http.HandlerFunc(notAllowed)

	router.PathPrefix("/").HandlerFunc(createSpaHandler(cfg.StaticDir, "index.html"))
	return router
}
