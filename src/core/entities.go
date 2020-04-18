package core

import (
	"audit/src/components/user"
	"audit/src/config"
)

func initEntities(cfg *config.AppConfig) {
	user.Init()
}
