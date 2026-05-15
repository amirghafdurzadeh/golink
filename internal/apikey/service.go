package apikey

import (
	"crypto/subtle"
	"errors"
)

var (
	ErrInvalidAPIKey = errors.New("invalid api key")
)

type Service interface {
	Validate(apiKey string) error
}

type service struct {
	expectedKey string
}

func NewService(expectedKey string) Service {
	return &service{
		expectedKey: expectedKey,
	}
}

func (s *service) Validate(apiKey string) error {
	if apiKey == "" {
		return ErrInvalidAPIKey
	}

	if subtle.ConstantTimeCompare([]byte(apiKey), []byte(s.expectedKey)) != 1 {
		return ErrInvalidAPIKey
	}

	return nil
}
