package http

import (
	"audit/src/auth/controller"
	"audit/src/context"
	"audit/src/utils"
	"net/http"

	"github.com/mitchellh/mapstructure"
)

// SignInRequestModel signin DTO
type SignInRequestModel struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// SignIn login handler
func SignIn(res http.ResponseWriter, req *http.Request) {
	var model SignInRequestModel
	err := mapstructure.Decode(context.GetContext(req).JSON(), &model)

	if err != nil {
		utils.ToError(res, 400, err)
		return
	}

	err = utils.ValidateModel(model)
	if err != nil {
		utils.ToError(res, 400, err)
		return
	}

	user, err := controller.SignIn(model.Username, model.Password)
	if err != nil {
		utils.ToError(res, 400, err)
		return
	}

	utils.ToJSON(res, http.StatusOK, user)
}
