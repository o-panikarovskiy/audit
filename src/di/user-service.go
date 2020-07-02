package di

import (
	"audit/src/user"
	"fmt"
)

// GetUserService return user.IService
func GetUserService() user.IService {
	var v user.IService
	ok := Get().Get(&v)

	if !ok {
		panic(fmt.Errorf("%T value not found in service locator", v))
	}

	return v
}
