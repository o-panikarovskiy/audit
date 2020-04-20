package handlers

import (
	"audit/src/di"
	"audit/src/middlewares"
	"audit/src/utils"
	"audit/src/utils/res"
	"net/http"

	"github.com/mitchellh/mapstructure"
)

type signInModel struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func validateSignInModel(r *http.Request) (*signInModel, error) {
	var model signInModel
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

// SignIn login handler
func SignIn(w http.ResponseWriter, r *http.Request) {
	model, err := validateSignInModel(r)
	if err != nil {
		res.ToError(w, http.StatusBadRequest, err, "INVALID_REQUEST_MODEL")
		return
	}

	user, err := di.GetUserService().Auth(model.Username, model.Password)
	if err != nil {
		res.ToError(w, http.StatusBadRequest, err)
		return
	}

	res.ToJSON(w, http.StatusOK, user)
}
