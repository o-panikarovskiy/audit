package middlewares

import (
	"context"
	"log"
	"net/http"
)

type ctxKeyType int

const (
	jsonKey ctxKeyType = iota
	sessionKey
	sessionUserKey
	// ...
)

// Context for middlewares
type Context struct {
	context.Context
}

// NewContext constructor
func NewContext(ctx context.Context) Context {
	return Context{ctx}
}

// MdlwTypedContext middleware for put Context to request
func MdlwTypedContext(next http.Handler) http.Handler {

	fn := func(res http.ResponseWriter, req *http.Request) {
		ctx := NewContext(req.Context())
		req = req.WithContext(ctx)
		next.ServeHTTP(res, req)
	}

	return http.HandlerFunc(fn)
}

// GetContext returns typed Context from request
func GetContext(r *http.Request) Context {
	ctx, ok := r.Context().(Context)
	if !ok {
		log.Panic("Failed to get custom Context from request")
	}

	return ctx
}
