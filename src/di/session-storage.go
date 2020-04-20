package di

import (
	"audit/src/sessions"
	"fmt"
)

// GetSessionStorage returns sessions.IStorage
func GetSessionStorage() sessions.IStorage {
	var v sessions.IStorage
	ok := Get().Get(&v)

	if !ok {
		panic(fmt.Errorf("%T value not found in service locator", v))
	}

	return v
}
