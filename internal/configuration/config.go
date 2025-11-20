package configuration

import (
	"os"
)

type Config struct {
	AppConfig   AppConfiguration
	BasicConfig BasicConfiguration
}

type Configurer interface {
	load()
}

func Init() Config {
	var appConfig AppConfiguration
	appConfig.load()
	var basicConfig BasicConfiguration
	basicConfig.load()
	config := Config{
		AppConfig:   appConfig,
		BasicConfig: basicConfig,
	}

	return config
}

func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value != "" {
		return value
	}
	return defaultValue
}
