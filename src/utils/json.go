package utils

import (
	"encoding/json"
	"io"
)

// JSONStringify func
func JSONStringify(data interface{}) (string, error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// JSONStringifyWriter func
func JSONStringifyWriter(w io.Writer, data interface{}) error {
	return json.NewEncoder(w).Encode(data)
}

// JSONParse func
func JSONParse(str string) (*map[string]interface{}, error) {
	dest := make(map[string]interface{})
	err := json.Unmarshal([]byte(str), &dest)

	if err != nil {
		return nil, err
	}

	return &dest, nil
}

// JSONParseReader func
func JSONParseReader(r io.Reader) (*map[string]interface{}, error) {
	dest := make(map[string]interface{})
	err := json.NewDecoder(r).Decode(&dest)

	if err != nil {
		return nil, err
	}

	return &dest, nil
}
