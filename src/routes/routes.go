package routes

import (
	"net/http"

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
	api.HandleFunc("/health", health)
	api.HandleFunc("/ws", sockets.SocketUpgradeHandler)

	api.Use(middlewares.ErrorHandle)
	api.NotFoundHandler = http.HandlerFunc(notFound)

	router.PathPrefix("/").HandlerFunc(createSpaHandler(cfg.StaticDir, "index.html"))

	return router
}

func health(w http.ResponseWriter, r *http.Request) {
	utils.Send(w, utils.StringMap{"ok": true})
}

func notFound(w http.ResponseWriter, r *http.Request) {
	utils.SendError(w, utils.NewAPIError(http.StatusNotFound, `NOT_FOUND`, `Path not found.`))
}
