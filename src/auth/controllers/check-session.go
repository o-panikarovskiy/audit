package controllers

import (
	"audit/src/di"
	"audit/src/user"
)

// CheckSession func
func CheckSession(sid string) (*user.User, error) {
	return di.GetUserService().CheckAuthSession(sid)
}
