package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"audit/src/config"
	"audit/src/routes"
)

var httpServer *http.Server

func createHTTPServer(cfg *config.AppConfig) *http.Server {
	wt := time.Second * 15
	rt := time.Second * 15
	it := time.Second * 60
	if cfg.IsDev() {
		wt = 0
		rt = 0
		it = 0
	}

	srv := &http.Server{
		Addr: fmt.Sprintf("0.0.0.0:%d", cfg.Port),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout:      wt,
		IdleTimeout:       it,
		ReadHeaderTimeout: rt,
		Handler:           routes.CreateRouter(cfg),
	}

	return srv
}

func shutdownHTTPServer(srv *http.Server, cfg *config.AppConfig) {
	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), cfg.GracefulTimeout*time.Second)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
}

func runHTTPServer(srv *http.Server) {
	log.Println(fmt.Sprintf("Server start listening on %v", srv.Addr))
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
