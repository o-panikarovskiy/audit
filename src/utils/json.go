package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

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

// DecodeJSONBody read json to dest
func DecodeJSONBody(r *http.Request, dest interface{}) error {
	if value := r.Header.Get("Content-Type"); !strings.HasPrefix(value, "application/json") {
		return BadRequestModel("Content-Type header is not application/json")
	}

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&dest)
	if err != nil {
		return BadRequestModel(err)
	}

	if dec.More() {
		return BadRequestModel("Request body must only contain a single JSON object")
	}

	return nil
}
