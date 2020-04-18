package http

import (
	"audit/src/components/auth/controller"
	"audit/src/utils"
	"net/http"
)

// SignUpRequestModel signup DTO
type SignUpRequestModel struct {
	Email    string `json:"email" validate:"required,email,max=256"`
	Password string `json:"password" validate:"required,min=8,max=256"`
}

// SignUp handler for create new user
func SignUp(w http.ResponseWriter, r *http.Request) {
	var model SignUpRequestModel
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

	user, err := controller.SignUp(model.Email, model.Password)
	if err != nil {
		utils.SendError(w, err)
		return
	}

	utils.Send(w, user)
}
