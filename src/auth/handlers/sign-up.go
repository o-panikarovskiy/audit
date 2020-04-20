package handlers

import (
	"audit/src/di"
	"audit/src/middlewares"
	"audit/src/utils"
	"audit/src/utils/res"
	"net/http"

	"github.com/mitchellh/mapstructure"
)

type signUpModel struct {
	Email    string `json:"email" validate:"required,email,max=256"`
	Password string `json:"password" validate:"required,min=8,max=256"`
}

func validateSignUpModel(r *http.Request) (*signUpModel, error) {
	var model signUpModel
	err := mapstructure.Decode(middlewares.GetContext(r).JSON(), &model)

	if err != nil {
		return nil, err
	}

	err = utils.ValidateModel(model)
	if err != nil {
		return nil, err
	}

	return &model, nil
}

// SignUp user register handler
func SignUp(w http.ResponseWriter, r *http.Request) {
	model, err := validateSignUpModel(r)
	if err != nil {
		res.ToError(w, http.StatusBadRequest, err, "INVALID_REQUEST_MODEL")
		return
	}

	user, err := di.GetUserService().Register(model.Email, model.Password)
	if err != nil {
		res.ToError(w, http.StatusBadRequest, err)
		return
	}

	res.ToJSON(w, http.StatusOK, user)
}
