package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"
)

// AppConfig main app config
type AppConfig struct {
	Port              int           `json:"port"`
	Env               string        `json:"env"`
	GracefulTimeout   time.Duration `json:"gracefulTimeout"`
	StaticDir         string        `json:"staticDir"`
	LogRequestAfterMs time.Duration `json:"logRequestAfterMs"`
	SessionAgeMin     int           `json:"SessionAgeMin"`
	Cookie            struct {
		Hash  string `json:"hash"`
		Block string `json:"block"`
	} `json:"cookie"`
	Redis struct {
		ConnectionString string `json:"connectionString"`
	} `json:"redis"`
	PG struct {
		ConnectionString     string `json:"connectionString"`
		PoolMinConns         int    `json:"poolMinConns"`
		PoolMaxConns         int    `json:"poolMaxConns"`
		MaxConnLifetimeMin   int    `json:"maxConnLifetimeMin"`
		MaxConnIdleTimeMin   int    `json:"maxConnIdleTimeMin"`
		HealthCheckPeriodMin int    `json:"healthCheckPeriodMin"`
	} `json:"pg"`
}

const (
	// DevMode indicates mode is debug.
	DevMode = "dev"
	// ProdMode indicates mode is production.
	ProdMode = "prod"
	// TestMode indicates mode is test.
	TestMode = "test"
)

// IsDev returns true if env in DevMode
func (c *AppConfig) IsDev() bool {
	return c.Env == DevMode
}

// IsProd returns true if env in ProdMode
func (c *AppConfig) IsProd() bool {
	return c.Env == ProdMode
}

// IsTest returns true if env in TestMode
func (c *AppConfig) IsTest() bool {
	return c.Env == TestMode
}

// NewDefaultConfig parses command line arguments and read json config
func NewDefaultConfig(path string) *AppConfig {
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
