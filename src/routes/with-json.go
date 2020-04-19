package routes

import (
	"audit/src/context"
	"audit/src/utils"
	"encoding/json"
	"net/http"
	"strings"
)

// WithJSON try to parse json from req.Body
func WithJSON(next http.Handler) http.Handler {
	fn := func(res http.ResponseWriter, req *http.Request) {
		data, err := decodeJSON(req)
		if err != nil {
			panic(err)
		}

		ctx := context.GetContext(req)
		ctx = ctx.WithJSON(data)
		req = req.WithContext(ctx)
		next.ServeHTTP(res, req)
	}

	return http.HandlerFunc(fn)
}

func decodeJSON(req *http.Request) (*utils.StringMap, error) {
	ct := req.Header.Get("Content-Type")
	if !strings.HasPrefix(ct, "application/json") {
		return nil, utils.BadRequestModel("Content-Type header is not application/json")
	}

	dest := make(utils.StringMap)
	dec := json.NewDecoder(req.Body)
	err := dec.Decode(&dest)

	if err != nil {
		return nil, utils.BadRequestModel(err)
	}

	return &dest, nil
}
