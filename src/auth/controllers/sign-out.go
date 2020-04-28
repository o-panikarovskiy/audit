package controllers

import (
	"audit/src/di"
	"audit/src/user"
)

// SignOut func
func SignOut(u *user.User) error {
	return di.GetUserService().SignOut(u)
}
