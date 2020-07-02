package main

import (
	"log"
	"os"
	"os/signal"
	"path/filepath"

	"audit/src/config"
	"audit/src/server"
)

func main() {
	inst := server.NewInstance(setup())

	inst.Run()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt) // SIGINT (Ctrl+C)
	<-c                            // Block until we receive our signal.

	inst.Stop()
	os.Exit(0)
}

func setup() *config.AppConfig {
	if len(os.Args) < 2 {
		log.Panicln("Please, specify the config file")
	}

	path, err := filepath.Abs(os.Args[1])
	if err != nil {
		log.Panicln(err)
	}

	return config.NewDefaultConfig(path)
}
