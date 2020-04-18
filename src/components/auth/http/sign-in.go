package http

import (
	"audit/src/components/auth/controller"
	"audit/src/utils"
	"net/http"
)

// SignInRequestModel signin DTO
type SignInRequestModel struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// SignIn handler for check auth
func SignIn(w http.ResponseWriter, r *http.Request) {
	var model SignInRequestModel
	err := utils.DecodeJSONBody(r, &model)

	if err != nil {
		utils.SendError(w, err)
		return
	}

	err = utils.ValidateModel(model)
	if err != nil {
		utils.SendError(w, err)
		return
	}

	user, err := controller.SignIn(model.Username, model.Password)
	if err != nil {
		utils.SendError(w, err)
		return
	}

	utils.Send(w, user)
}
