package controllers

import (
	"audit/src/di"
	"audit/src/user"
	"audit/src/utils"

	"github.com/mitchellh/mapstructure"
)

// SignUpReqModel struct
type SignUpReqModel struct {
	Email    string `json:"email" validate:"required,email,max=256"`
	Password string `json:"password" validate:"required,min=8,max=256"`
}

// SignUpResModel struct
type SignUpResModel struct {
	User *user.User `json:"user"`
	SID  string     `json:"sid" `
}

// SignUp func
func SignUp(model *SignUpReqModel) (*SignUpResModel, error) {
	user, sid, err := di.GetUserService().Register(model.Email, model.Password)
	if err != nil {
		return nil, err
	}

	return &SignUpResModel{User: user, SID: sid}, nil
}

// ValidateSignUp func
func ValidateSignUp(json *utils.StringMap) (*SignUpReqModel, error) {
	var model SignUpReqModel
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
