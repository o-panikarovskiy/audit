package config

import (
	"audit/src/utils"
	"log"
	"os"
	"path/filepath"

	"github.com/mitchellh/mapstructure"
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
