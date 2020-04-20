package controller

import (
	"audit/src/di"
	"audit/src/middlewares"
	"audit/src/utils"
	"audit/src/utils/res"
	"net/http"

	"github.com/mitchellh/mapstructure"
)

// SignInRequestModel signin DTO
type SignInRequestModel struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// SignIn login handler
func SignIn(w http.ResponseWriter, r *http.Request) {
	var model SignInRequestModel
	err := mapstructure.Decode(middlewares.GetContext(r).JSON(), &model)

	if err != nil {
		res.ToError(w, http.StatusBadRequest, err, "INVALID_REQUEST_MODEL")
		return
	}

	err = utils.ValidateModel(model)
	if err != nil {
		res.ToError(w, http.StatusBadRequest, err)
		return
	}

	service := di.Get().GetUserService()
	user, err := service.Auth(model.Username, model.Password)
	if err != nil {
		res.ToError(w, http.StatusBadRequest, err)
		return
	}

	res.ToJSON(w, http.StatusOK, user)
}
