package handlers

import (
	"audit/src/auth/controllers"
	"audit/src/utils/res"
	"net/http"

	"github.com/gorilla/mux"
)

// EndRegistration handler
func EndRegistration(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tokenID := vars["token"]

	model, err := controllers.EndSignUp(tokenID)
	if err != nil {
		res.SendStatusError(w, http.StatusBadRequest, err)
		return
	}

	setAuthCookie(w, model.SID)
	res.SendJSON(w, http.StatusOK, *model.User)
}
