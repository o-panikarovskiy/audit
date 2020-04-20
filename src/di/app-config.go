package di

import "audit/src/config"

// GetAppConfig return *config.AppConfig
func GetAppConfig() *config.AppConfig {
	var cfg *config.AppConfig
	Get().Get(&cfg)
	return cfg
}
