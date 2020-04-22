package server

import (
	"audit/src/config"
	"audit/src/di"
	"audit/src/sessions/redisses"
	"audit/src/user/defservice"
	"audit/src/user/pgrep"
)

func createDevInstase(cfg *config.AppConfig) *Instance {
	pool, err := createPgxPool(cfg)
	if err != nil {
		panic(err)
	}

	rep, err := pgrep.NewRepository(pool)
	if err != nil {
		panic(err)
	}

	ses, err := redisses.NewStorage((cfg))
	if err != nil {
		panic(err)
	}

	deps := &di.ServiceLocator{}
	deps.Register(cfg)
	deps.Register(ses)
	deps.Register(defservice.NewDefaultUserService(rep, ses, cfg))

	di.Set(deps)

	addSocketEventListeners(cfg)
	return &Instance{
		cfg:        cfg,
		httpServer: createHTTPServer(cfg),
	}
}
