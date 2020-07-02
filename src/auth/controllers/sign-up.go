package controllers

import (
	"audit/src/di"
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
	TokenID string `json:"sid" `
}

// SignUp func
func SignUp(json *map[string]interface{}) (*SignUpResModel, error) {
	model, err := validateSignUp(json)
	if err != nil {
		return nil, err
	}

	tokenID, _, err := di.GetUserService().SignUp(model.Email, model.Password)
	if err != nil {
		return nil, err
	}

	return &SignUpResModel{TokenID: tokenID}, nil
}

func validateSignUp(json *map[string]interface{}) (*SignUpReqModel, error) {
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
