package utils

import (
	"audit/src/config"
	"encoding/json"
	"net/http"
)

// ToJSON send json to response
func ToJSON(res http.ResponseWriter, status int, data ...interface{}) {
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.WriteHeader(status)
	json.NewEncoder(res).Encode(data)
}

// ToError sends error
func ToError(res http.ResponseWriter, status int, err error) {
	var appErr *AppError
	cfg := config.GetCurrentConfig()

	if e, ok := err.(*AppError); ok {
		appErr = e
	} else {
		appErr = &AppError{
			Status:  status,
			Code:    "APP_ERROR",
			Message: err.Error(),
			Err:     err,
		}
	}

	appErr.WriteJSON(res, status, cfg.IsDev())
}
