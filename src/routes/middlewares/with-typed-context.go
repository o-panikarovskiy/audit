package middlewares

import (
	"net/http"
)

// WithTypedContext middleware for returning Context
func WithTypedContext(next http.Handler) http.Handler {

	fn := func(res http.ResponseWriter, req *http.Request) {
		ctx := NewContext(req.Context())
		req = req.WithContext(ctx)
		next.ServeHTTP(res, req)
	}

	return http.HandlerFunc(fn)
}
