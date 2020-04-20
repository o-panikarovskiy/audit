package middlewares

import (
	"audit/src/config"
	"net/http"
)

// WithAppConfig middleware
func WithAppConfig(next http.Handler) http.Handler {

	fn := func(res http.ResponseWriter, req *http.Request) {
		ctx := GetContext(req)
		ctx = ctx.WithAppConfig(config.GetCurrentConfig())
		req = req.WithContext(ctx)
		next.ServeHTTP(res, req)
	}

	return http.HandlerFunc(fn)
}
