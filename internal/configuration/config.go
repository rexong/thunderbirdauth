package configuration

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AppConfig   AppConfiguration
	BasicConfig BasicConfiguration
	LdapConfig  LdapConfiguration
}

type Configurer interface {
	load()
}

func Init() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error Loading .env file: %v", err)
	}
	var appConfig AppConfiguration
	appConfig.load()
	var basicConfig BasicConfiguration
	basicConfig.load()
	var ldapConfig LdapConfiguration
	ldapConfig.load()

	config := Config{
		AppConfig:   appConfig,
		BasicConfig: basicConfig,
		LdapConfig:  ldapConfig,
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

func getBoolEnv(key string, defaultBool bool) bool {
	value := os.Getenv(key)
	if value == "" {
		return defaultBool
	}
	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		return defaultBool
	}
	return boolValue
}
