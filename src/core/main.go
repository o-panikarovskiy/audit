package core

import (
	"audit/src/config"
	"fmt"
	"log"
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

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		log.Println(fmt.Sprintf("Server start listening on %d port", inst.cfg.Port))
		if err := inst.httpServer.ListenAndServe(); err != nil {
			panic(err)
		}
	}()
}

// Stop instanse
func (inst *Instance) Stop() {
	shutDownEntities(inst.cfg)
	shutdownHTTPServer(inst.httpServer, inst.cfg)
}
