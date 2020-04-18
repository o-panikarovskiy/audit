package http

import (
	"audit/src/components/auth/controller"
	"audit/src/utils"
	"net/http"
)

// Login handler
func Login(w http.ResponseWriter, r *http.Request) {
	user, err := controller.Login("oleg.pnk@gmail.com", "123")

	if err != nil {
		utils.SendError(w, err)
		return
	}

	utils.Send(w, user)
}
