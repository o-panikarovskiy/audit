package userservice

import (
	"audit/src/user"
	"audit/src/utils"

	"github.com/mitchellh/mapstructure"
)

func (s *userService) CheckAuthSession(sid string) (*user.User, error) {
	sessionUser, err := s.RestoreSessionUser(sid)
	if err != nil {
		return nil, &utils.AppError{Code: "APP_ERROR", Message: err.Error(), Err: err}
	}

	if sessionUser == nil {
		return nil, &utils.AppError{Code: "AUTH_ERROR", Message: "User not found"}
	}

	dbUser, err := s.FindByID(sessionUser.ID)
	if err != nil {
		return nil, &utils.AppError{Code: "APP_ERROR", Message: err.Error(), Err: err}
	}

	s.saveAuthSession(sid, dbUser) // update session age
	return dbUser, err
}

func (s *userService) RestoreSessionUser(sid string) (*user.User, error) {
	var target user.User

	json, err := s.sessions.GetJSON(authUserKey + sid)
	if err != nil || json == nil {
		return nil, err
	}

	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		DecodeHook: utils.StringToDateTimeHook,
		Result:     &target,
	})

	if err != nil {
		return nil, err
	}

	err = decoder.Decode(json)
	if err != nil {
		return nil, err
	}

	return &target, nil
}
