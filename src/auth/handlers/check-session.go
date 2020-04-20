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
	sessionUser := middlewares.GetContext(r).GetSessionUser()

	user, err := controllers.CheckSession(sid, sessionUser)

	if err != nil {
		res.SendStatusError(w, http.StatusUnauthorized, err)
		return
	}

	res.SendJSON(w, http.StatusOK, user)
}
