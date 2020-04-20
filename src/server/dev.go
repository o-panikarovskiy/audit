package server

import (
	"audit/src/config"
	"audit/src/di"
	"audit/src/user"
)

func createDevInstase(cfg *config.AppConfig) *Instance {
	addSocketEventListeners(cfg)

	deps := &di.ServiceLocator{}
	rep := user.NewTestRepository()

	deps.Register(cfg)
	deps.Register(user.NewUserService(rep))

	di.Set(deps)

	return &Instance{
		cfg:        cfg,
		httpServer: createHTTPServer(cfg),
	}

}
