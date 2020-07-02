package handlers

import (
	"audit/src/auth/controllers"
	"audit/src/middlewares"
	"audit/src/utils/res"
	"net/http"
)

// SignIn login handler
func SignIn(w http.ResponseWriter, r *http.Request) {
	model, err := controllers.SignIn(middlewares.GetContext(r).JSON())

	if err != nil {
		res.SendStatusError(w, http.StatusBadRequest, err)
		return
	}

	setAuthCookie(w, model.SID)
	res.SendJSON(w, http.StatusOK, *model.User)
}
