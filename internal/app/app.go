package app

import (
	"context"
	"errors"
	"time"

	"github.com/amirghafdurzadeh/golink/internal/apikey"
	"github.com/amirghafdurzadeh/golink/internal/config"
	"github.com/amirghafdurzadeh/golink/internal/database"
	"github.com/amirghafdurzadeh/golink/internal/health"
	"github.com/amirghafdurzadeh/golink/internal/link"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type Application interface {
	Config() *config.Config
	Services() Services
	Close() error
}

type application struct {
	cfg      *config.Config
	postgres *pgxpool.Pool
	redis    *redis.Client
	services Services
}

func New(ctx context.Context) (Application, error) {
	// config
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	// infrastructure
	postgres, err := database.NewPostgres(ctx, cfg.PostgresURL)
	if err != nil {
		return nil, err
	}

	redis, err := database.NewRedis(ctx, cfg.RedisAddr, cfg.RedisPassword)
	if err != nil {
		postgres.Close()
		return nil, err
	}

	// repositories
	linkRepo := link.NewPostgresRepository(postgres)

	// caches
	linkCache := link.NewRedisCache(redis, 24*time.Hour)

	// services
	apikeyService := apikey.NewService(cfg.APIKey)
	healthService := health.NewService(postgres, redis)
	linkService := link.NewService(linkRepo, linkCache, cfg.ShortCodeLength)

	return &application{
		cfg:      cfg,
		postgres: postgres,
		redis:    redis,
		services: NewServices(
			apikeyService,
			healthService,
			linkService,
		),
	}, nil
}

func (a *application) Config() *config.Config {
	return a.cfg
}

func (a *application) Services() Services {
	return a.services
}

func (a *application) Close() error {
	var errs []error

	if a.postgres != nil {
		a.postgres.Close()
	}

	if a.redis != nil {
		if err := a.redis.Close(); err != nil {
			errs = append(errs, err)
		}
	}

	return errors.Join(errs...)
}
