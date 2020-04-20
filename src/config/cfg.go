package config

import (
	"time"
)

// AppConfig main app config
type AppConfig struct {
	Port              int           `json:"port"`
	Env               string        `json:"env"`
	GracefulTimeout   time.Duration `json:"gracefulTimeout"`
	StaticDir         string        `json:"staticDir"`
	LogRequestAfterMs time.Duration `json:"logRequestAfterMs"`
	SessionAge        int           `json:"sessionAge"`
	Cookie            struct {
		Hash  string `json:"hash"`
		Block string `json:"block"`
	} `json:"cookie"`
	Redis struct {
		Port int    `json:"port"`
		Host string `json:"host"`
	} `json:"redis"`
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
