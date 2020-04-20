package routes

import (
	"net/http"

	"audit/src/auth"
	"audit/src/config"
	"audit/src/routes/middlewares"
	"audit/src/sockets"

	"github.com/gorilla/mux"
)

// CreateRouter create main http.Handler
func CreateRouter(cfg *config.AppConfig) http.Handler {
	router := mux.NewRouter()

	api := router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/health", health).Methods("GET")
	api.HandleFunc("/ws", sockets.HTTPUpgradeHandler).Methods("GET")

	api.Use(middlewares.WithError)
	api.Use(middlewares.WithTypedContext)
	api.Use(middlewares.WithAppConfig)
	api.Use(middlewares.WithLog)

	auth.GetRoutes(api)

	api.NotFoundHandler = http.HandlerFunc(notFound)
	api.MethodNotAllowedHandler = http.HandlerFunc(notAllowed)

	router.PathPrefix("/").HandlerFunc(createSpaHandler(cfg.StaticDir, "index.html"))
	return router
}
