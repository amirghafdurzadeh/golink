package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Environment     string
	AppPort         string
	PostgresURL     string
	RedisAddr       string
	RedisPassword   string
	APIKey          string
	ShortCodeLength int
}

func Load() (*Config, error) {
	appEnv := getEnv("APP_ENV", "development")
	if appEnv != "production" {
		err := godotenv.Load()
		if err != nil {
			return nil, err
		}
	}

	postgresHost := getEnv("POSTGRES_HOST", "localhost")
	postgresPort := getEnv("POSTGRES_PORT", "5432")
	postgresUser := getEnv("POSTGRES_USER", "postgres")
	postgresPass := getEnv("POSTGRES_PASSWORD", "postgres")
	postgresDB := getEnv("POSTGRES_DB", "golink")

	postgresURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		postgresUser, postgresPass, postgresHost, postgresPort, postgresDB,
	)

	redisHost := getEnv("REDIS_HOST", "localhost")
	redisPort := getEnv("REDIS_PORT", "6379")
	redisAddr := fmt.Sprintf("%s:%s", redisHost, redisPort)

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		return nil, errAPIKeyRequired
	}

	shortCodeLength, err := getEnvAsInt("SHORT_CODE_LENGTH", 6)
	if err != nil {
		return nil, err
	}

	return &Config{
		Environment:     appEnv,
		AppPort:         getEnv("APP_PORT", "8080"),
		PostgresURL:     postgresURL,
		RedisAddr:       redisAddr,
		RedisPassword:   getEnv("REDIS_PASSWORD", ""),
		APIKey:          apiKey,
		ShortCodeLength: shortCodeLength,
	}, nil
}
