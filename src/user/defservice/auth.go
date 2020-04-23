package defservice

import (
	"audit/src/user"
	"audit/src/utils"
	"strings"
)

const authSidKey = "AUTH:SID:"
const authUserKey = "AUTH:USER:"

func (s *userService) Auth(email string, password string) (*user.User, string, error) {
	email = strings.ToLower(email)
	user, err := s.FindByUsername(email)

	if err != nil {
		return nil, "", &utils.AppError{Code: "APP_ERROR", Message: err.Error(), Err: err}
	}

	if user == nil ||
		user.PasswordHash != utils.SHA512(password, user.PasswordSalt) {
		return nil, "", &utils.AppError{Code: "AUTH_ERROR", Message: "Email or password is incorrect"}
	}

	err = s.destroyAuthSession(user.ID)
	if err != nil {
		return nil, "", &utils.AppError{Code: "APP_ERROR", Message: err.Error(), Err: err}
	}

	sid := utils.RandomString(64)
	err = s.saveAuthSession(sid, user)
	if err != nil {
		return nil, "", &utils.AppError{Code: "APP_ERROR", Message: err.Error(), Err: err}
	}

	return user, sid, nil
}

func (s *userService) saveAuthSession(sid string, user *user.User) error {
	expiration := s.cfg.SessionAgeSec

	err := s.sessions.SetJSON(authUserKey+sid, user, expiration)
	if err != nil {
		return err
	}

	err = s.sessions.Set(authSidKey+user.ID, sid, expiration)
	if err != nil {
		return err
	}

	return nil
}

func (s *userService) destroyAuthSession(userID string) error {
	sid, err := s.sessions.Get(authSidKey + userID)

	if err != nil {
		return err
	}

	if sid != "" {
		_, err = s.sessions.Delete(sid)
	}

	if err != nil {
		return err
	}

	return nil
}
