package userservice

import (
	"audit/src/user"
	"audit/src/utils"
	"strings"
)

func (s *userService) Auth(email string, password string) (*user.User, string, error) {
	if email == "" || password == "" {
		return nil, "", invalidReqModelErr
	}

	email = strings.ToLower(email)
	user, err := s.FindByUsername(email)
	if err != nil {
		return nil, "", err
	}

	if user == nil ||
		user.PasswordHash != utils.SHA512(password, user.PasswordSalt) {
		return nil, "", authAppErr
	}

	err = s.destroyAuthSession(user.ID)
	if err != nil {
		return nil, "", err
	}

	sid := utils.RandomString(64)
	err = s.saveAuthSession(sid, user)
	if err != nil {
		return nil, "", err
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
		_, err = s.sessions.Delete(authUserKey + sid)
	}

	if err != nil {
		return err
	}

	_, err = s.sessions.Delete(authSidKey + userID)
	if err != nil {
		return err
	}

	return nil
}
