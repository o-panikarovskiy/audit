package di

import (
	"audit/src/config"
	"audit/src/user"
)

// Injector struct
type Injector struct {
	cfg         *config.AppConfig
	userService user.IService
}

var deps *Injector

// New create default DI
func New(
	cfg *config.AppConfig,
	userService user.IService,
) *Injector {

	deps = &Injector{
		cfg:         cfg,
		userService: userService,
	}

	return deps
}

// Get returns current DI
func Get() *Injector { return deps }

// GetAppConfig return service
func (di *Injector) GetAppConfig() *config.AppConfig { return di.cfg }

// GetUserService return service
func (di *Injector) GetUserService() user.IService { return di.userService }
