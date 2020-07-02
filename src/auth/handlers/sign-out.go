package handlers

import (
	"audit/src/auth/controllers"
	"audit/src/middlewares"
	"audit/src/utils/res"
	"net/http"
)

// SignOut handler
func SignOut(w http.ResponseWriter, r *http.Request) {
	user := middlewares.GetContext(r).GetSessionUser()
	controllers.SignOut(user)

	res.SendNoContent(w)
}
