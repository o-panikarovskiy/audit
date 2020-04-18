package utils

import (
	"encoding/json"
	"net/http"
)

// Send send data with 200 http status code
func Send(w http.ResponseWriter, data interface{}) {
	SendJSON(w, http.StatusOK, data)
}

// SendJSON send json to response
func SendJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// SendError send json to response
func SendError(w http.ResponseWriter, err error) {
	if e, ok := err.(*APPError); ok {
		SendJSON(w, e.Status, e)
		return
	}
	SendJSON(w, http.StatusInternalServerError, err)
}
