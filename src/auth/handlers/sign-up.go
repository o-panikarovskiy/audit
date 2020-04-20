package handlers

import (
	"audit/src/auth/controllers"
	"audit/src/middlewares"
	"audit/src/utils/res"
	"net/http"
)

// SignUp user register handler
func SignUp(w http.ResponseWriter, r *http.Request) {
	model, err := controllers.SignUp(middlewares.GetContext(r).JSON())

	if err != nil {
		res.SendStatusError(w, http.StatusBadRequest, err)
		return
	}

	SetCookie(w, model.SID)
	res.SendJSON(w, http.StatusOK, &model.User)
}
