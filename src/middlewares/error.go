package middlewares

import (
	"log"
	"net/http"

	"audit/src/config"
	"audit/src/utils"
)

// ErrorHandle global handle error
func ErrorHandle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			r := recover()
			if r == nil {
				return
			}

			var err error
			switch t := r.(type) {
			case string:
				err = utils.NewAPIError(http.StatusInternalServerError, "APP_ERROR", t)
			case error:
				err = t
			default:
				err = utils.NewAPIError(http.StatusInternalServerError, "APP_ERROR", "Unknown error")
			}

			cfg := config.GetCurrentConfig()
			apiErr, ok := err.(*utils.APIError)
			if ok && cfg.Env == config.ProdMode {
				apiErr.Stack = nil
			}

			log.Println("api error: ", err)

			utils.SendError(w, err)
		}()

		next.ServeHTTP(w, r)
	})
}
