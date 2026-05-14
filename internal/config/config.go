package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort         string
	PostgresConnURL string
	RedisAddr       string
	RedisPassword   string
	APIKey          string
}

func Load() (*Config, error) {
	if os.Getenv("APP_ENV") != "production" {
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

	postgresConnURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		postgresUser, postgresPass, postgresHost, postgresPort, postgresDB,
	)

	redisHost := getEnv("REDIS_HOST", "localhost")
	redisPort := getEnv("REDIS_PORT", "6379")
	redisAddr := fmt.Sprintf("%s:%s", redisHost, redisPort)

	return &Config{
		AppPort:         getEnv("APP_PORT", "8080"),
		PostgresConnURL: postgresConnURL,
		RedisAddr:       redisAddr,
		RedisPassword:   getEnv("REDIS_PASSWORD", ""),
		APIKey:          getEnv("API_KEY", "my_secret_apikey"),
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
