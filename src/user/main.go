package user

import (
	"audit/src/config"
)

// Init db coonection and services
func Init(cfg *config.AppConfig) IService {
	var repository IRepository

	if cfg.IsDev() || cfg.IsTest() {
		repository = NewTestRepository()
	} else {
		repository = NewPostgreSQLRepository("some pool")
	}

	return NewUserService(repository)
}

// ShutDown allows grasefull exit
func ShutDown(srv IService, cfg *config.AppConfig) {
	srv.ShutDown()
}
