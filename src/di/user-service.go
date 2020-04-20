package di

import (
	"audit/src/user"
)

// GetUserService return user.IService
func GetUserService() user.IService {
	var s user.IService
	Get().Get(&s)
	return s
}
