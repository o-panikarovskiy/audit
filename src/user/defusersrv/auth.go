package defusersrv

import (
	"audit/src/user"
	"audit/src/utils"
	"strings"
)

func (s *userService) Auth(email string, password string) (*user.User, string, error) {
	email = strings.ToLower(email)
	user, err := s.FindByEmail(email)

	if err != nil {
		return nil, "", utils.NewAppError("APP_ERROR", err.Error())
	}

	if user == nil ||
		user.PasswordHash != utils.SHA512(password, user.PasswordSalt) {
		return nil, "", utils.NewAppError("AUTH_ERROR", "Email or password is incorrect")
	}

	err = s.destroyAuthSession(user.ID)
	if err != nil {
		return nil, "", utils.NewAppError("APP_ERROR", err.Error())
	}

	sid, err := s.createAuthSession(user)
	if err != nil {
		return nil, "", utils.NewAppError("APP_ERROR", err.Error())
	}

	return user, sid, nil
}
