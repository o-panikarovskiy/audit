package controller

import "audit/src/components/user"

// SignUp create new user
func SignUp(email string, password string) (*user.User, error) {
	service := user.GetService()
	return service.Register(email, password)
}
