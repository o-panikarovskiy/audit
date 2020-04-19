package routes

import (
	"log"
	"net/http"

	"audit/src/config"
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

			switch val := r.(type) {
			case *utils.AppError:
				var cfg = config.GetCurrentConfig()
				if cfg.IsProd() {
					val.Stack = nil
				}
				utils.SendJSON(w, val.Status, val)
			case error:
				utils.SendJSON(w, http.StatusInternalServerError, &utils.StringMap{"message": val.Error()})
			case string:
				utils.SendJSON(w, http.StatusInternalServerError, &utils.StringMap{"message": val})
			default:
				http.Error(w, "Unknown error", http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
