package server

import (
	"audit/src/config"
	"audit/src/di"
	"audit/src/user"
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
	rep := user.NewTestRepository()

	di.New(
		inst.cfg,
		user.NewUserStore(rep),
	)

	addSocketEventListeners(inst.cfg)
	go runHTTPServer(inst.httpServer)
}

// Stop instanse
func (inst *Instance) Stop() {
	shutdownHTTPServer(inst.httpServer, inst.cfg)
}
