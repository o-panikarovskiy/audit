package di

import (
	"audit/src/config"
)

// Injector struct
type Injector struct {
}

var deps *Injector

// New create default DI
func New(cfg *config.AppConfig) *Injector {
	deps = &Injector{}

	return deps
}

// GetInjector returns current DI
func GetInjector() *Injector {
	return deps
}
