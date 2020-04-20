package di

import (
	"audit/src/config"
	"fmt"
)

// GetAppConfig return *config.AppConfig
func GetAppConfig() *config.AppConfig {
	var v *config.AppConfig
	ok := Get().Get(&v)

	if !ok {
		panic(fmt.Errorf("%T value not found in service locator", v))
	}

	return v
}
