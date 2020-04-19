package controller

import "audit/src/user"

// CheckSession checks user session by id
func CheckSession(sessionID string) (*user.User, error) {
	service := user.GetService()
	return service.CheckSession(sessionID)
}
