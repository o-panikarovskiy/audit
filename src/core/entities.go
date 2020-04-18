package core

import (
	"audit/src/components/user"
	"audit/src/config"
)

func initEntities(cfg *config.AppConfig) {
	user.Init(cfg)
}

func shutDownEntities(cfg *config.AppConfig) {
	user.ShutDown(cfg)
}
