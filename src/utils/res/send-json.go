package res

import (
	"encoding/json"
	"net/http"
)

// SendJSON send json to http response
func SendJSON(w http.ResponseWriter, status int, data interface{}) error {
	WriteJSONHeader(w, status)
	return json.NewEncoder(w).Encode(data)
}

// SendNoContent send 204 wuth application/json header
func SendNoContent(w http.ResponseWriter) {
	WriteJSONHeader(w, http.StatusNoContent)
}

// WriteJSONHeader func
func WriteJSONHeader(w http.ResponseWriter, status int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
}
