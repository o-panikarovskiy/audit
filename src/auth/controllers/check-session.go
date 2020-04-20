package controllers

import (
	"audit/src/di"
	"audit/src/user"
	"fmt"
)

// CheckSession func
func CheckSession(sid string, sessionUser *map[string]interface{}) (*user.User, error) {
	id := fmt.Sprint((*sessionUser)["ID"])
	return di.GetUserService().CheckSession(sid, id)
}
