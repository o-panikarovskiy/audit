package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// NewDefaultConfig parses command line arguments and read json config
func NewDefaultConfig() *AppConfig {
	if len(os.Args) < 2 {
		log.Panicln("Please, specify the config file")
	}

	path, err := filepath.Abs(os.Args[1])
	if err != nil {
		log.Panicln(err)
	}

	return readConfigFile(path)
}

func readConfigFile(path string) *AppConfig {
	jsonFile, err := os.Open(path)
	if err != nil {
		log.Panicln(err)
	}

	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Panicln(err)
	}

	var result AppConfig
	err = json.Unmarshal(byteValue, &result)
	if err != nil {
		log.Panicln(err)
	}

	return &result
}
