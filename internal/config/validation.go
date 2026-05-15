package config

import "errors"

var (
	errAPIKeyIsRequired  = errors.New("API_KEY is required")
	errBaseURLIsRequired = errors.New("BASE_URL is required")
)

func (c *Config) Validate() error {
	if c.App.APIKey == "" {
		return errAPIKeyIsRequired
	}

	if c.App.BaseURL == "" {
		return errBaseURLIsRequired
	}

	return nil
}
