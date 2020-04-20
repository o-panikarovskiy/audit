package server

import (
	"audit/src/config"
	"audit/src/user"
)

func initEntities(cfg *config.AppConfig) {

	user.Init(cfg)
}

func shutDownEntities(cfg *config.AppConfig) {

	//	user.ShutDown(cfg)
}
