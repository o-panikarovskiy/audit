package controller

import "audit/src/components/user"

// SignIn return exist user
func SignIn(username string, password string) (*user.User, error) {
	service := user.GetService()
	return service.Auth(username, password)
}
