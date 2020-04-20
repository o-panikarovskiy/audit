package controller

import (
	"net/http"
)

// SignInRequestModel signin DTO
type SignInRequestModel struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// SignIn login handler
func SignIn(w http.ResponseWriter, r *http.Request) {
	// var model SignInRequestModel
	// err := mapstructure.Decode(injector.GetContext(r).JSON(), &model)

	// if err != nil {
	// 	res.ToError(w, http.StatusBadRequest, err, "INVALID_REQUEST_MODEL")
	// 	return
	// }

	// err = utils.ValidateModel(model)
	// if err != nil {
	// 	res.ToError(w, http.StatusBadRequest, err)
	// 	return
	// }

	// user, err := controller.SignIn(model.Username, model.Password)
	// if err != nil {
	// 	res.ToError(w, http.StatusBadRequest, err)
	// 	return
	// }

	// res.ToJSON(w, http.StatusOK, user)
}
