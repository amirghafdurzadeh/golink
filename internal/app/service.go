package app

import (
	"github.com/amirghafdurzadeh/golink/internal/apikey"
	"github.com/amirghafdurzadeh/golink/internal/health"
	"github.com/amirghafdurzadeh/golink/internal/link"
)

type Services interface {
	APIKey() apikey.Service
	Health() health.Service
	Link() link.Service
}

type services struct {
	apiKeyService apikey.Service
	healthService health.Service
	linkService   link.Service
}

func NewServices(
	apiKeyService apikey.Service,
	healthService health.Service,
	linkService link.Service,
) Services {

	return &services{
		apiKeyService: apiKeyService,
		healthService: healthService,
		linkService:   linkService,
	}
}

func (s *services) APIKey() apikey.Service {
	return s.apiKeyService
}

func (s *services) Health() health.Service {
	return s.healthService
}

func (s *services) Link() link.Service {
	return s.linkService
}
