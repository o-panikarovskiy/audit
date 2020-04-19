package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"audit/src/config"
	"audit/src/routes"
)

var httpServer *http.Server

func createHTTPServer(cfg *config.AppConfig) *http.Server {
	srv := &http.Server{
		Addr: fmt.Sprintf("0.0.0.0:%d", cfg.Port),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout:      time.Second * 15,
		IdleTimeout:       time.Second * 60,
		ReadHeaderTimeout: time.Second * 15,
		Handler:           routes.CreateRouter(cfg),
	}

	return srv
}

func shutdownHTTPServer(srv *http.Server, cfg *config.AppConfig) {
	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), cfg.GracefulTimeout)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
}
