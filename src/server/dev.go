package server

import (
	"audit/src/config"
	"audit/src/di"
	"audit/src/sessions/redisses"
	"audit/src/user/defservice"
	"audit/src/user/emailconfirm"
	"audit/src/user/pgrep"
)

func createDevInstase(cfg *config.AppConfig) *Instance {
	pgPool, err := createPgxPool(cfg)
	if err != nil {
		panic(err)
	}

	pgRepository, err := pgrep.NewRepository(pgPool)
	if err != nil {
		panic(err)
	}

	redisStorage, err := redisses.NewStorage((cfg))
	if err != nil {
		panic(err)
	}

	emailConfirmator := emailconfirm.NewEmailConfirmService(cfg)
	userService := defservice.NewDefaultUserService(pgRepository, redisStorage, emailConfirmator, cfg)

	deps := &di.ServiceLocator{}
	deps.Register(cfg)
	deps.Register(redisStorage)
	deps.Register(userService)

	di.Set(deps)

	addSocketEventListeners(cfg)
	return &Instance{
		cfg:        cfg,
		httpServer: createHTTPServer(cfg),
	}
}
