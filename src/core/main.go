package core

import (
	"audit/src/config"
)

// Run servers
func Run(cfg *config.AppConfig) {
	initEntities(cfg)
	addSocketEventListeners(cfg)

	createHTTPServer(cfg)
}

// Stop servers
func Stop(cfg *config.AppConfig) {
	shutDownEntities(cfg)
	shutdownHTTPServer(cfg)
}
