package routes

import (
	"net/http"

	"github.com/o-panikarovskiy/audit/src/middlewares"
	"github.com/o-panikarovskiy/audit/src/sockets"
	"github.com/o-panikarovskiy/audit/src/utils"

	"github.com/gorilla/mux"
)

// CreateRouter create main routes
func CreateRouter() *mux.Router {
	router := mux.NewRouter()

	api := router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/health", health)
	api.HandleFunc("/ws", sockets.SocketUpgradeHandler)

	api.Use(middlewares.ErrorHandle)
	api.NotFoundHandler = http.HandlerFunc(notFound)

	router.PathPrefix("/").HandlerFunc(createSpaHandler("client/dist", "index.html"))

	addSocketEventListeners()

	return router
}

func health(w http.ResponseWriter, r *http.Request) {
	utils.Send(w, utils.HT{"ok": true})
}

func notFound(w http.ResponseWriter, r *http.Request) {
	utils.SendError(w, utils.NewAPIError(http.StatusNotFound, `NOT_FOUND`, `Path not found.`))
}
