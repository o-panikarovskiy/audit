package middlewares

import (
	"audit/src/di"
	"audit/src/utils/res"
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/securecookie"
)

// MdlwSession check user session from sid cookie
// and set user
func MdlwSession(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		cfg := di.GetAppConfig()

		hashKey := []byte(cfg.Cookie.Hash)
		blockKey := []byte(cfg.Cookie.Block)
		s := securecookie.New(hashKey, blockKey)

		cookie, err := r.Cookie("sid")
		if err != nil {
			res.SendStatusError(w, http.StatusUnauthorized, err, "SESSION_ERROR")
			return
		}

		sid := ""
		err = s.Decode("sid", cookie.Value, &sid)
		if err != nil {
			res.SendStatusError(w, http.StatusUnauthorized, err, "SESSION_ERROR")
			return
		}

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
	val, ok := ctx.Value(sessionKey).(string)

	if !ok {
		panic(fmt.Errorf("Failed to get value from context %v by key %v", val, sessionKey))
	}

	return val
}
