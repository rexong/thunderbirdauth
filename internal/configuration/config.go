package configuration

import (
	"os"
)

type Config struct {
	AppConfig AppConfiguration
}

type Configurer interface {
	load()
}

func Init() Config {
	var appConfig AppConfiguration
	appConfig.load()
	config := Config{AppConfig: appConfig}

	return config
}

func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value != "" {
		return value
	}
	return defaultValue
}
