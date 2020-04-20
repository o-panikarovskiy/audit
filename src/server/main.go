package server

import (
	"audit/src/config"
	"audit/src/di"
	"net/http"
)

// Instance server
type Instance struct {
	httpServer *http.Server
	cfg        *config.AppConfig
}

// NewInstance create new Instance
func NewInstance(cfg *config.AppConfig) *Instance {
	if cfg.IsDev() {
		return createDevInstase(cfg)
	}
	return nil
}

// Run instanse
func (inst *Instance) Run() {
	go runHTTPServer(inst.httpServer)
}

// Stop instanse
func (inst *Instance) Stop() {
	di.GetUserService().ShutDown()
	shutdownHTTPServer(inst.httpServer, inst.cfg)
}
