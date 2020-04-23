package middlewares

import (
	"audit/src/di"
	"audit/src/user"
	"audit/src/utils/res"
	"context"
	"fmt"
	"net/http"
)

// MdlwSessionUser set user from sid
func MdlwSessionUser(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		service := di.GetUserService()

		sid := GetContext(r).GetSessionID()
		su, err := service.RestoreSessionUser(sid)
		if err != nil {
			res.SendStatusError(w, http.StatusUnauthorized, err, "SESSION_ERROR")
			return
		}

		ctx := GetContext(r)
		ctx = ctx.WithSessionUser(su)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

// WithSessionUser put session user to context
func (ctx Context) WithSessionUser(su *user.User) Context {
	return Context{context.WithValue(ctx, sessionUserKey, su)}
}

// GetSessionUser get session user from context
func (ctx Context) GetSessionUser() *user.User {
	val, ok := ctx.Value(sessionUserKey).(*user.User)

	if !ok {
		panic(fmt.Errorf("Failed to get value from context %v by key %v", val, sessionUserKey))
	}

	return val
}
