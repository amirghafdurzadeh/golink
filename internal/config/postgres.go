package config

import "fmt"

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DB       string
	SSLMode  string
}

func (c PostgresConfig) URL() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.DB,
		c.SSLMode,
	)
}

// for logging
func (c PostgresConfig) SafeURL() string {
	return fmt.Sprintf(
		"postgres://%s:***@%s:%s/%s?sslmode=%s",
		c.User,
		c.Host,
		c.Port,
		c.DB,
		c.SSLMode,
	)
}
