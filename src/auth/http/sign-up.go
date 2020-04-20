package http

import (
	"audit/src/auth/controller"
	"audit/src/routes/middlewares"
	"audit/src/utils"
	"audit/src/utils/res"
	"net/http"

	"github.com/mitchellh/mapstructure"
)

// SignUpRequestModel signup DTO
type SignUpRequestModel struct {
	Email    string `json:"email" validate:"required,email,max=256"`
	Password string `json:"password" validate:"required,min=8,max=256"`
}

// SignUp handler for create new user
func SignUp(w http.ResponseWriter, r *http.Request) {
	var model SignUpRequestModel
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

	user, err := controller.SignUp(model.Email, model.Password)
	if err != nil {
		res.ToError(w, http.StatusBadRequest, err)
		return
	}

	res.ToJSON(w, http.StatusOK, user)
}
