package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// AppConfigCookie struct
type AppConfigCookie struct {
	Hash  string `json:"hash"`
	Block string `json:"block"`
	Name  string `json:"name"`
}

// AppConfigRedis struct
type AppConfigRedis struct {
	ConnectionString string `json:"connectionString"`
}

// AppConfigPG struct
type AppConfigPG struct {
	ConnectionString     string `json:"connectionString"`
	PoolMinConns         int    `json:"poolMinConns"`
	PoolMaxConns         int    `json:"poolMaxConns"`
	MaxConnLifetimeMin   int    `json:"maxConnLifetimeMin"`
	MaxConnIdleTimeMin   int    `json:"maxConnIdleTimeMin"`
	HealthCheckPeriodMin int    `json:"healthCheckPeriodMin"`
}

// AppConfigRateLimit struct
type AppConfigRateLimit struct {
	IntervalMs  int `json:"intervalMs"`
	MaxRequests int `json:"maxRequests"`
}

// AppConfig main app config
type AppConfig struct {
	Port               int                `json:"port"`
	Env                string             `json:"env"`
	GracefulTimeoutSec int                `json:"gracefulTimeoutSec"`
	StaticDir          string             `json:"staticDir"`
	LogRequestAfterMs  int                `json:"logRequestAfterMs"`
	SessionAgeSec      int                `json:"sessionAgeSec"`
	Cookie             AppConfigCookie    `json:"cookie"`
	RateLimit          AppConfigRateLimit `json:"rateLimit"`
	Redis              AppConfigRedis     `json:"redis"`
	PG                 AppConfigPG        `json:"pg"`
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
