package userservice

import (
	"audit/src/user"
)

func (s *userService) SignOut(u *user.User) error {
	if u == nil {
		return nil
	}

	err := s.destroyAuthSession(u.ID)
	if err != nil {
		return err
	}

	return nil
}
