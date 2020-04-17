package config

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"audit/src/utils"

	"github.com/mitchellh/mapstructure"
)

// AppConfig main app config
type AppConfig struct {
	Port            int           `json:"port"`
	Env             string        `json:"env"`
	GracefulTimeout time.Duration `json:"gracefulTimeout"`
	StaticDir       string        `json:"staticDir"`
}

const (
	// DevMode indicates mode is debug.
	DevMode = "dev"
	// ProdMode indicates mode is production.
	ProdMode = "prod"
	// TestMode indicates mode is test.
	TestMode = "test"
)

var currentConfig *AppConfig = nil

// GetCurrentConfig return current config
func GetCurrentConfig() *AppConfig {
	return currentConfig
}

// ReadConfig parses command line arguments and read json config
func ReadConfig() *AppConfig {
	if len(os.Args) < 2 {
		log.Panicln("Please, specify the config file")
	}

	path, err := filepath.Abs(os.Args[1])
	if err != nil {
		log.Panicln(err)
	}

	c, err := utils.ReadJSONFile(path)
	if err != nil {
		log.Panicln(err)
	}

	err = mapstructure.Decode(c, &currentConfig)
	if err != nil {
		log.Panicln(err)
	}

	return currentConfig
}
