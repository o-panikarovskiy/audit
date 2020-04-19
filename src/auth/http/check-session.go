package http

import (
	"audit/src/auth/controller"
	"audit/src/utils"
	"net/http"
)

// CheckSession handler
func CheckSession(w http.ResponseWriter, r *http.Request) {
	user, err := controller.CheckSession("test")

	if err != nil {
		panic(err)
	}

	utils.SendJSON(w, http.StatusOK, user)
}
