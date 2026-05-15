package config

type Config struct {
	HTTP     HTTPConfig
	Redis    RedisConfig
	Postgres PostgresConfig
	App      AppConfig
}

func Load() (*Config, error) {
	cfg := &Config{
		HTTP: HTTPConfig{
			Port:              getEnv("HTTP_PORT", "8080"),
			ReadTimeout:       getEnvAsDuration("HTTP_READ_TIMEOUT", "5s"),
			ReadHeaderTimeout: getEnvAsDuration("HTTP_READ_HEADER_TIMEOUT", "5s"),
			WriteTimeout:      getEnvAsDuration("HTTP_WRITE_TIMEOUT", "10s"),
			IdleTimeout:       getEnvAsDuration("HTTP_IDLE_TIMEOUT", "60s"),
		},
		Redis: RedisConfig{
			Host:     getEnv("REDIS_HOST", "localhost"),
			Port:     getEnv("REDIS_PORT", "6379"),
			Password: getEnv("REDIS_PASSWORD", ""),
		},
		Postgres: PostgresConfig{
			Host:     getEnv("POSTGRES_HOST", "localhost"),
			Port:     getEnv("POSTGRES_PORT", "5432"),
			User:     getEnv("POSTGRES_USER", "postgres"),
			Password: getEnv("POSTGRES_PASSWORD", "postgres"),
			DB:       getEnv("POSTGRES_DB", "golink"),
			SSLMode:  getEnv("POSTGRES_SSL_MODE", "disable"),
		},
		App: AppConfig{
			Environment:     getEnv("APP_ENV", "development"),
			APIKey:          getEnv("API_KEY", ""),
			ShortCodeLength: getEnvAsInt("SHORT_CODE_LENGTH", 6),
			BaseURL:         getEnv("BASE_URL", ""),
		},
	}

	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}
