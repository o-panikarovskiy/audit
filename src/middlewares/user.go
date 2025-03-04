package middlewares

import (
	"audit/src/di"
	"audit/src/user"
	"audit/src/utils"
	"audit/src/utils/res"
	"context"
	"net/http"
)

// MdlwSessionUser set user from sid
func MdlwSessionUser(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		service := di.GetUserService()

		sid := GetContext(r).GetSessionID()
		if sid == "" {
			res.SendStatusError(w, http.StatusUnauthorized, &utils.AppError{Code: "SESSION_ERROR"})
			return
		}

		su, err := service.RestoreSessionUser(sid)
		if err != nil {
			res.SendStatusError(w, http.StatusUnauthorized, &utils.AppError{Code: "SESSION_ERROR", Err: err})
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
	raw := ctx.Value(sessionUserKey)
	if raw == nil {
		return nil
	}

	val, ok := raw.(*user.User)
	if !ok {
		return nil
	}

	return val
}
