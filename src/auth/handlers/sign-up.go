package handlers

import (
	"audit/src/auth/controllers"
	"audit/src/middlewares"
	"audit/src/utils/res"
	"net/http"
)

// SignUp user register handler
func SignUp(w http.ResponseWriter, r *http.Request) {
	reqModel, err := controllers.ValidateSignUp(middlewares.GetContext(r).JSON())
	if err != nil {
		res.ToError(w, http.StatusBadRequest, err, "INVALID_REQUEST_MODEL")
		return
	}

	resModel, err := controllers.SignUp(reqModel)
	if err != nil {
		res.ToError(w, http.StatusBadRequest, err)
		return
	}

	res.ToJSON(w, http.StatusOK, &resModel)
}
