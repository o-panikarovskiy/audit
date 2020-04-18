package http

import (
	"audit/src/components/auth/controller"
	"audit/src/utils"
	"net/http"
)

// CheckSession handler
func CheckSession(w http.ResponseWriter, r *http.Request) {
	user, err := controller.CheckSession("test")

	if err != nil {
		utils.SendError(w, err)
		return
	}

	utils.Send(w, user)
}
