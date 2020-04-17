package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/o-panikarovskiy/audit/src/config"
	"github.com/o-panikarovskiy/audit/src/server"
)

func main() {
	cfg := config.GetCurrentConfig()
	httpServer := server.StartHTTPServer(cfg)

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	server.ShutdownHTTPServer(httpServer, cfg.GracefulTimeout)

	log.Println("shutting down")
	os.Exit(0)
}
