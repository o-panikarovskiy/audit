package defusersrv

import (
	"audit/src/user"
	"audit/src/utils"
)

func (s *userService) CheckSession(sid string, userID string) (*user.User, error) {
	user, err := s.Find(userID)

	if err != nil {
		return nil, &utils.AppError{Code: "AUTH_ERROR", Message: err.Error(), Err: err}
	}

	if user == nil {
		return nil, &utils.AppError{Code: "AUTH_ERROR", Message: "User not found"}
	}

	s.saveAuthSession(sid, user)

	return user, err
}
