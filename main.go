package main

import (
	"log"
	"os"
	"os/signal"

	"audit/src/config"
	"audit/src/server"
)

func main() {
	cfg := config.ReadConfig()
	server.Run(cfg)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt) // SIGINT (Ctrl+C)
	<-c                            // Block until we receive our signal.

	server.Stop(cfg)
	log.Println("shutting down")
	os.Exit(0)
}
