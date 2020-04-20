package middlewares

import (
	"audit/src/utils"
	"audit/src/utils/res"
	"encoding/json"
	"net/http"
	"strings"
)

// WithJSON try to parse json from req.Body
func WithJSON(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		data, err := decodeJSON(r)
		if err != nil {
			res.ToError(w, http.StatusBadRequest, err)
			return
		}

		ctx := GetContext(r)
		ctx = ctx.WithJSON(data)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func decodeJSON(req *http.Request) (*utils.StringMap, error) {
	ct := req.Header.Get("Content-Type")
	if !strings.HasPrefix(ct, "application/json") {
		return nil, utils.NewAppError("INVALID_HEADERS", "Content-Type header is not application/json")
	}

	dest := make(utils.StringMap)
	dec := json.NewDecoder(req.Body)
	err := dec.Decode(&dest)

	if err != nil {
		return nil, err
	}

	return &dest, nil
}
