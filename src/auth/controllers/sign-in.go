package controllers

import (
	"audit/src/di"
	"audit/src/user"
	"audit/src/utils"

	"github.com/mitchellh/mapstructure"
)

// SignInReqModel model
type SignInReqModel struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// SignInResModel struct
type SignInResModel struct {
	User *user.User `json:"user"`
	SID  string     `json:"sid" `
}

// SignIn func
func SignIn(model *SignInReqModel) (*SignInResModel, error) {
	user, sid, err := di.GetUserService().Auth(model.Username, model.Password)
	if err != nil {
		return nil, err
	}

	return &SignInResModel{User: user, SID: sid}, nil
}

// ValidateSignIn func
func ValidateSignIn(json *utils.StringMap) (*SignInReqModel, error) {
	var model SignInReqModel
	err := mapstructure.Decode(json, &model)

	if err != nil {
		return nil, err
	}

	err = utils.ValidateModel(model)
	if err != nil {
		return nil, err
	}

	return &model, nil
}
