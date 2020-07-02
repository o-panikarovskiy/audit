package handlers

import (
	"audit/src/auth/controllers"
	"audit/src/middlewares"
	"audit/src/utils/res"
	"net/http"
)

// CheckSession handler
func CheckSession(w http.ResponseWriter, r *http.Request) {
	sid := middlewares.GetContext(r).GetSessionID()
	user, err := controllers.CheckSession(sid)

	if err != nil {
		res.SendStatusError(w, http.StatusUnauthorized, err)
		return
	}

	setAuthCookie(w, sid) // update cookie age
	res.SendJSON(w, http.StatusOK, *user)
}
