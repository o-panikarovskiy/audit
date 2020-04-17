package config

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/o-panikarovskiy/audit/src/utils"

	"github.com/mitchellh/mapstructure"
)

// AppConfig main app config
type AppConfig struct {
	Port            int
	Env             string
	GracefulTimeout time.Duration
}

const (
	// DevMode indicates gin mode is debug.
	DevMode = "dev"
	// ProdMode indicates gin mode is release.
	ProdMode = "prod"
	// TestMode indicates gin mode is test.
	TestMode = "test"
)

var currentConfig *AppConfig = nil

// GetCurrentConfig return current config
func GetCurrentConfig() *AppConfig {
	if currentConfig == nil {
		currentConfig = ReadConfig()
	}
	return currentConfig
}

// ReadConfig parses command line arguments and read json config
func ReadConfig() *AppConfig {
	mode := getEnv()
	c, err := utils.ReadJSONFile(getJSONFileName(mode))

	if err != nil {
		panic(err)
	}

	var result AppConfig
	err = mapstructure.Decode(c, &result)

	if err != nil {
		panic(err)
	}

	return &result
}

func getEnv() string {
	mode := os.Getenv("APP_ENV")

	if mode == "" {
		mode = DevMode
	}

	mode = *flag.String("mode", mode, "dev | prod | test")
	flag.Parse()

	return mode
}

func getJSONFileName(mode string) string {
	curdir, err := os.Getwd()
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%s/src/config/%s.json", curdir, mode)
}
