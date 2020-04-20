package defusersrv

import (
	"audit/src/user"
	"audit/src/utils"
)

func (s *userService) createAuthSession(user *user.User) (string, error) {
	sid := utils.RandomString(64)
	return s.saveAuthSession(sid, user)
}
