package context

import (
	"log"
	"net/http"
)

// GetContext returns typed context from request
func GetContext(r *http.Request) Context {
	ctx, ok := r.Context().(Context)
	if !ok {
		log.Panic("Failed to get custom Context from request")
		return ctx
	}
	return ctx
}
