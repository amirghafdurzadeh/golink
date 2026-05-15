package config

import (
	"os"
	"strconv"
	"time"
)

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	value := getEnv(key, strconv.Itoa(defaultValue))

	n, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}

	return n
}

func getEnvAsDuration(key, defaultValue string) time.Duration {
	value := getEnv(key, defaultValue)

	d, err := time.ParseDuration(value)
	if err != nil {
		d, _ = time.ParseDuration(defaultValue)
	}

	return d
}
