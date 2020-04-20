package server

import (
	"audit/src/config"
	"audit/src/di"
	"audit/src/sessions"
	"audit/src/user"
)

func createDevInstase(cfg *config.AppConfig) *Instance {
	addSocketEventListeners(cfg)

	deps := &di.ServiceLocator{}
	rep := user.NewTestRepository()
	ses, err := sessions.NewRedisStorage((cfg))
	if err != nil {
		panic(err)
	}

	deps.Register(cfg)
	deps.Register(ses)
	deps.Register(user.NewUserService(rep, ses, cfg))

	di.Set(deps)

	return &Instance{
		cfg:        cfg,
		httpServer: createHTTPServer(cfg),
	}
}
