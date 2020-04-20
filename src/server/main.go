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
	inst := &Instance{
		cfg:        cfg,
		httpServer: createHTTPServer(cfg),
	}

	return inst
}

// Run instanse
func (inst *Instance) Run() {
	initEntities(inst.cfg)
	addSocketEventListeners(inst.cfg)

	di.New(inst.cfg)

	// Run our server in a goroutine so that it doesn't block.
	go runHTTPServer(inst.httpServer)
}

// Stop instanse
func (inst *Instance) Stop() {
	shutDownEntities(inst.cfg)
	shutdownHTTPServer(inst.httpServer, inst.cfg)
}
