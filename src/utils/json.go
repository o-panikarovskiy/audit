package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

// StringMap is a shortcut for map[string]interface{}
type StringMap map[string]interface{}

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

// ReadJSONFile read json file to map
func ReadJSONFile(path string) (*StringMap, error) {
	jsonFile, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		return nil, err
	}

	var result StringMap
	err = json.Unmarshal(byteValue, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}
