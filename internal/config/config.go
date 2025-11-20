package config

import (
	"os"
)

type Configurer interface {
	load()
}

func Init() AppConfiguration {
	var appConfig AppConfiguration
	appConfig.load()
	return appConfig
}

func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value != "" {
		return value
	}
	return defaultValue
}
