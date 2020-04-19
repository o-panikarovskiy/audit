package routes

import (
	"audit/src/config"
	"log"
	"net/http"
	"time"
)

// WithLog create log middleware
func WithLog(next http.Handler) http.Handler {
	fn := func(res http.ResponseWriter, req *http.Request) {
		start := time.Now()

		defer func() {
			cfg := config.GetCurrentConfig()
			maxElapsed := cfg.LogRequestAfterMs * time.Millisecond

			if d := time.Since(start); d > maxElapsed {
				log.Printf("%v %s %s", d, req.Method, req.URL)
			}
		}()

		next.ServeHTTP(res, req)
	}
	return http.HandlerFunc(fn)
}
