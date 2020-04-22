package defusersrv

import (
	"audit/src/user"
	"audit/src/utils"
	"strings"
)

func (s *userService) Register(email string, password string) (*user.User, string, error) {
	email = strings.ToLower(email)

	exUser, err := s.FindByEmail(email)
	if err != nil {
		return nil, "", utils.NewAppError("APP_ERROR", err.Error())
	}

	if exUser != nil {
		return nil, "", utils.NewAppError("USER_EXISTS", "User already exists")
	}

	salt := utils.RandomString(64)
	user := &user.User{
		Email:        email,
		Role:         user.UserRole,
		PasswordSalt: salt,
		PasswordHash: utils.SHA512(password, salt),
	}

	dbuser, err := s.Store(user)
	if err != nil {
		return nil, "", utils.NewAppError("APP_ERROR", err.Error())
	}

	sid, err := s.createAuthSession(dbuser)
	if err != nil {
		return nil, "", utils.NewAppError("APP_ERROR", err.Error())
	}

	return dbuser, sid, nil
}
