package config

import (
	"os"
	"strconv"
)

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) (int, error) {
	value := getEnv(key, strconv.Itoa(defaultValue))
	return strconv.Atoi(value)
}
