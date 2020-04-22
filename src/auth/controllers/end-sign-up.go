package controllers

import (
	"audit/src/di"
	"audit/src/user"
)

// EndSignUpResModel struct
type EndSignUpResModel struct {
	User *user.User `json:"user"`
	SID  string     `json:"sid" `
}

// EndSignUp func
func EndSignUp(tokenID string) (*EndSignUpResModel, error) {
	user, sid, err := di.GetUserService().EndSignUp(tokenID)
	if err != nil {
		return nil, err
	}

	return &EndSignUpResModel{User: user, SID: sid}, nil
}
