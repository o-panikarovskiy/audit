package middlewares

import (
	"audit/src/utils"
	"audit/src/utils/res"
	"context"
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
	raw := ctx.Value(jsonKey)
	if raw == nil {
		return emptyJSON()
	}

	val, ok := raw.(*map[string]interface{})
	if !ok {
		return emptyJSON()
	}

	return val
}

func emptyJSON() *map[string]interface{} {
	json := make(map[string]interface{})
	return &json
}

func decodeJSON(r *http.Request) (*map[string]interface{}, error) {
	ct := r.Header.Get("Content-Type")
	if !strings.HasPrefix(ct, "application/json") {
		return nil, utils.NewAppError("INVALID_HEADERS", "Content-Type header is not application/json")
	}

	return utils.JSONParseReader(r.Body)
}
