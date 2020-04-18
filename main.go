package main

import (
	"os"
	"os/signal"

	"audit/src/config"
	"audit/src/core"
)

func main() {
	cfg := config.ReadConfig()
	core.Run(cfg)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt) // SIGINT (Ctrl+C)
	<-c                            // Block until we receive our signal.

	core.Stop(cfg)
	os.Exit(0)
}
