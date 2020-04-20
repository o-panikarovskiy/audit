package defusersrv

import (
	"audit/src/user"
	"audit/src/utils"
)

func (s *userService) Register(email string, password string) (*user.User, string, error) {
	exUser, err := s.FindByEmail(email)
	if err != nil {
		return nil, "", utils.NewAppError("APP_ERROR", err.Error())
	}

	if exUser != nil {
		return nil, "", utils.NewAppError("USER_EXISTS", "User already exists")
	}

	salt := utils.RandomString(64)
	user := &user.User{
		ID:           utils.CreateGUID(),
		Email:        email,
		Role:         user.UserRole,
		PasswordSalt: salt,
		PasswordHash: utils.SHA512(password, salt),
	}

	sid, err := s.createAuthSession(user)
	if err != nil {
		return nil, "", utils.NewAppError("APP_ERROR", err.Error())
	}

	s.Store(user)
	return user, sid, nil
}
