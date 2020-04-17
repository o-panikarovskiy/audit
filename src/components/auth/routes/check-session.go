package routes

import (
	"audit/src/components/auth/bl"
	"audit/src/utils"
	"net/http"
)

func checkSession(w http.ResponseWriter, r *http.Request) {
	user, err := bl.CheckSession("test")

	if err != nil {
		utils.SendError(w, err)
		return
	}

	utils.Send(w, user)
}
