package server

import (
	"audit/src/config"
	"audit/src/di"
	"audit/src/sessions/redisses"
	"audit/src/user/emailconfirmator"
	"audit/src/user/pgrep"
	"audit/src/user/userservice"
)

func createDevInstase(cfg *config.AppConfig) *Instance {
	pgDb, err := openPosgresDB(cfg.PG.ConnectionString)
	if err != nil {
		panic(err)
	}

	pgRepository, err := pgrep.NewRepository(pgDb)
	if err != nil {
		panic(err)
	}

	redisStorage, err := redisses.NewStorage(cfg)
	if err != nil {
		panic(err)
	}

	confirmator := emailconfirmator.NewEmailConfirmatorService(cfg)
	userService := userservice.NewDefaultUserService(pgRepository, redisStorage, confirmator, cfg)

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
