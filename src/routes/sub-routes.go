package routes

import (
	"audit/src/utils"

	"github.com/gorilla/mux"
)

func setSubRoutes(router *mux.Router, prefix string, routeHandlers utils.RouteHandlers) {
	sub := router.PathPrefix(prefix).Subrouter()
	for _, val := range *routeHandlers {
		sub.HandleFunc(val.Route, val.Handler).Methods(val.Method)
	}
}
