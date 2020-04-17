package server

import "audit/src/config"

// Run servers
func Run(cfg *config.AppConfig) {
	createHTTPServer(cfg)

	addSocketEventListeners()
}

// Stop servers
func Stop(cfg *config.AppConfig) {
	shutdownHTTPServer(cfg)
}
