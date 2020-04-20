package http

import (
	"audit/src/auth/controller"
	"audit/src/utils/res"
	"net/http"
)

// CheckSession handler
func CheckSession(w http.ResponseWriter, r *http.Request) {
	user, err := controller.CheckSession("test")

	if err != nil {
		res.ToError(w, http.StatusBadRequest, err)
		return
	}

	res.ToJSON(w, http.StatusOK, user)
}
