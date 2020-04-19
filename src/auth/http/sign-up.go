package http

import (
	"audit/src/auth/controller"
	"audit/src/context"
	"audit/src/utils"
	"net/http"

	"github.com/mitchellh/mapstructure"
)

// SignUpRequestModel signup DTO
type SignUpRequestModel struct {
	Email    string `json:"email" validate:"required,email,max=256"`
	Password string `json:"password" validate:"required,min=8,max=256"`
}

// SignUp handler for create new user
func SignUp(res http.ResponseWriter, req *http.Request) {
	var model SignUpRequestModel
	err := mapstructure.Decode(context.GetContext(req).JSON(), &model)

	if err != nil {
		panic(err)
	}

	err = utils.ValidateModel(model)
	if err != nil {
		panic(err)
	}

	user, err := controller.SignUp(model.Email, model.Password)
	if err != nil {
		panic(err)
	}

	utils.SendJSON(res, http.StatusOK, user)
}
