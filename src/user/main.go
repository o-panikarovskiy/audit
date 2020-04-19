package user

import "audit/src/config"

var service IService
var repository IRepository

// Init db coonection and services
func Init(cfg *config.AppConfig) {
	if cfg.IsDev() || cfg.IsTest() {
		repository = NewTestRepository()
	} else {
		repository = NewPostgreSQLRepository("some pool")
	}

	service = NewUserService(repository)
}

// GetService returns Service that implements business logic
func GetService() IService {
	return service
}

// ShutDown allows grasefull exit
func ShutDown(cfg *config.AppConfig) {
	repository.ShutDown()
}
