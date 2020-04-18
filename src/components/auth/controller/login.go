package controller

import "audit/src/components/user"

// Login checks user session by id
func Login(email string, password string) (*user.User, error) {
	service := user.GetService()
	return service.Auth(email, password)
}
