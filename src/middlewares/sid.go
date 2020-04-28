package middlewares

import (
	"audit/src/di"
	"audit/src/utils/res"
	"context"
	"net/http"

	"github.com/gorilla/securecookie"
)

// MdlwSession check user session from sid cookie
func MdlwSession(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		cfg := di.GetAppConfig()

		hashKey := []byte(cfg.Cookie.Hash)
		blockKey := []byte(cfg.Cookie.Block)
		s := securecookie.New(hashKey, blockKey)

		cookie, err := r.Cookie(cfg.Cookie.Name)
		if err != nil {
			res.SendStatusError(w, http.StatusUnauthorized, err, "SESSION_ERROR")
			return
		}

		sid := ""
		err = s.Decode(cfg.Cookie.Name, cookie.Value, &sid)
		if err != nil {
			res.SendStatusError(w, http.StatusUnauthorized, err, "SESSION_ERROR")
			return
		}

		ctx := GetContext(r)
		ctx = ctx.WithSessionID(sid)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

// WithSessionID put sid to context
func (ctx Context) WithSessionID(data string) Context {
	return Context{context.WithValue(ctx, sessionKey, data)}
}

// GetSessionID get sid from context
func (ctx Context) GetSessionID() string {
	raw := ctx.Value(sessionKey)
	if raw == nil {
		return ""
	}

	val, ok := raw.(string)
	if !ok {
		return ""
	}

	return val
}
