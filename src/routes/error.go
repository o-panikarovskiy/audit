package routes

import (
	"log"
	"net/http"

	"audit/src/utils"
)

// WithError global handle error
func WithError(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			r := recover()
			if r == nil {
				return
			}

			log.Println("api error: ", r)

			if err, ok := r.(error); ok {
				utils.ToError(w, http.StatusInternalServerError, err)
			} else {
				http.Error(w, "Unknown error", http.StatusInternalServerError)
			}

		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
