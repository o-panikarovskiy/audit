package main

import (
	"os"
	"os/signal"

	"audit/src/config"
	"audit/src/server"
)

func main() {
	inst := server.NewInstance(config.ReadConfig())
	inst.Run()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt) // SIGINT (Ctrl+C)
	<-c                            // Block until we receive our signal.

	inst.Stop()
	os.Exit(0)
}
