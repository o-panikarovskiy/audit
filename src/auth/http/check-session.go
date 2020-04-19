package http

import (
	"audit/src/auth/controller"
	"audit/src/utils"
	"net/http"
)

// CheckSession handler
func CheckSession(res http.ResponseWriter, req *http.Request) {
	user, err := controller.CheckSession("test")

	if err != nil {
		utils.ToError(res, 400, err)
		return
	}

	utils.ToJSON(res, http.StatusOK, user)
}
