package controllers

import (
	"audit/src/di"
	"audit/src/sockets"
	"audit/src/user"
)

// SignOut func
func SignOut(u *user.User) error {
	if u == nil {
		return nil
	}

	err := di.GetUserService().SignOut(u)
	if err != nil {
		return err
	}

	sc := sockets.FindClient(func(clientID string, userID string) bool {
		if userID == u.ID {
			return true
		}
		return false
	})

	if sc != nil {
		sockets.RemoveClient(sc.GetID())
	}

	return nil
}
