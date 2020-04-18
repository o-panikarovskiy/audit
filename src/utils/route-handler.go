package utils

import "net/http"

// RouteHandler helps build routes
type RouteHandler struct {
	Route   string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
}

// RouteHandlers is array of RouteHandler
type RouteHandlers *[]RouteHandler
