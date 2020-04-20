package res

import (
	"encoding/json"
	"net/http"
)

// SendJSON send json to http response
func SendJSON(res http.ResponseWriter, status int, data interface{}) error {
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.WriteHeader(status)
	return json.NewEncoder(res).Encode(data)
}
