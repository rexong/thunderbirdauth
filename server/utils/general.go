package utils

import (
	"log"
	"os"
)

func GetEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	log.Printf("%s: %s", key, value)
	if value != "" {
		return value
	}
	return defaultValue
}

func ShouldSeed() bool {
	return os.Getenv("ENV") != PROD_ENV
}
