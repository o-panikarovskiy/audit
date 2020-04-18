package routes

import (
	"net/http"

	"audit/src/components/auth"
	"audit/src/config"
	"audit/src/middlewares"
	"audit/src/sockets"
	"audit/src/utils"

	"github.com/gorilla/mux"
)

// CreateRouter create main http.Handler
func CreateRouter(cfg *config.AppConfig) http.Handler {
	router := mux.NewRouter()

	api := router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/health", health).Methods("GET")
	api.HandleFunc("/ws", sockets.HTTPUpgradeHandler).Methods("GET")

	setSubRoutes(api, "/auth", auth.GetRoutes())

	api.Use(middlewares.ErrorHandle)
	api.NotFoundHandler = http.HandlerFunc(notFound)

	router.PathPrefix("/").HandlerFunc(createSpaHandler(cfg.StaticDir, "index.html"))

	return router
}

func setSubRoutes(router *mux.Router, prefix string, routeHandlers utils.RouteHandlers) {
	sub := router.PathPrefix(prefix).Subrouter()
	for _, val := range *routeHandlers {
		sub.HandleFunc(val.Route, val.Handler).Methods(val.Method)
	}
}

func health(w http.ResponseWriter, r *http.Request) {
	utils.Send(w, utils.StringMap{"ok": true})
}

func notFound(w http.ResponseWriter, r *http.Request) {
	utils.SendError(w, utils.NewAPPError(http.StatusNotFound, `NOT_FOUND`, `Path not found.`))
}
