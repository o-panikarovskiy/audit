package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/o-panikarovskiy/audit/src/config"
	"github.com/o-panikarovskiy/audit/src/routes"
)

// StartHTTPServer start http server
func StartHTTPServer(cfg *config.AppConfig) *http.Server {
	srv := &http.Server{
		Addr: fmt.Sprintf("0.0.0.0:%d", cfg.Port),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout:      time.Second * 15,
		IdleTimeout:       time.Second * 60,
		ReadHeaderTimeout: time.Second * 15,
		Handler:           routes.CreateRouter(),
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		log.Println(fmt.Sprintf("Server start listening on 0.0.0.0:%d", cfg.Port))
		if err := srv.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	return srv
}

// ShutdownHTTPServer shutdown http server
func ShutdownHTTPServer(httpServer *http.Server, waitTime time.Duration) {
	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), waitTime)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	httpServer.Shutdown(ctx)
}
