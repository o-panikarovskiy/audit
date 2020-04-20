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
		res.ToError(w, http.StatusBadRequest, err)
		return
	}

	res.ToJSON(w, http.StatusOK, model)
}
