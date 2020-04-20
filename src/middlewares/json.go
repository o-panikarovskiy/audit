package middlewares

import (
	"audit/src/utils"
	"audit/src/utils/res"
	"context"
	"fmt"
	"net/http"
	"strings"
)

// MdlwJSON try to parse json from req.Body
func MdlwJSON(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		data, err := decodeJSON(r)
		if err != nil {
			res.SendStatusError(w, http.StatusBadRequest, err)
			return
		}

		ctx := GetContext(r)
		ctx = ctx.WithJSON(data)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

// WithJSON put json to context
func (ctx Context) WithJSON(data *map[string]interface{}) Context {
	return Context{context.WithValue(ctx, jsonKey, data)}
}

// JSON get json data from context
func (ctx Context) JSON() *map[string]interface{} {
	val, ok := ctx.Value(jsonKey).(*map[string]interface{})

	if !ok {
		panic(fmt.Errorf("Failed to get value from context %v by key %v", val, jsonKey))
	}

	return val
}

func decodeJSON(r *http.Request) (*map[string]interface{}, error) {
	ct := r.Header.Get("Content-Type")
	if !strings.HasPrefix(ct, "application/json") {
		return nil, utils.NewAppError("INVALID_HEADERS", "Content-Type header is not application/json")
	}

	return utils.JSONParseReader(r.Body)
}
