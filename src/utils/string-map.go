package utils

import "net/http"

// StringMap is a shortcut for map[string]interface{}
type StringMap map[string]interface{}

// RouteHandler helps build routes
type RouteHandler struct {
	Route   string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
}

// RouteHandlers is array of RouteHandler
type RouteHandlers *[]RouteHandler
