package defusersrv

import (
	"audit/src/user"
	"audit/src/utils"
)

func (s *userService) saveAuthSession(sid string, user *user.User) (string, error) {
	val, err := utils.JSONStringify(user)

	if err != nil {
		return "", err
	}

	err = s.sessions.Set(sid, val, s.cfg.SessionAge)
	if err != nil {
		return "", err
	}

	err = s.sessions.Set(user.ID, sid, s.cfg.SessionAge)
	if err != nil {
		return "", err
	}

	return sid, nil
}
