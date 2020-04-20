package middlewares

import (
	"audit/src/di"
	"log"
	"net/http"
	"time"
)

// MdlwLog create log middleware
func MdlwLog(next http.Handler) http.Handler {
	fn := func(res http.ResponseWriter, req *http.Request) {
		start := time.Now()

		defer func() {
			cfg := di.GetAppConfig()
			maxElapsed := cfg.LogRequestAfterMs * time.Millisecond

			if d := time.Since(start); d > maxElapsed {
				log.Printf("%v %s %s", d, req.Method, req.URL)
			}
		}()

		next.ServeHTTP(res, req)
	}
	return http.HandlerFunc(fn)
}
