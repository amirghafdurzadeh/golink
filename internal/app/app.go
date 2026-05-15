package app

import (
	"context"
	"errors"
	"sync"
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
	closeOnce sync.Once
	cfg       *config.Config
	postgres  *pgxpool.Pool
	redis     *redis.Client
	services  Services
}

func New(ctx context.Context) (Application, error) {
	// config
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	// infrastructure
	postgres, err := database.NewPostgres(ctx, cfg.Postgres.URL())
	if err != nil {
		return nil, err
	}

	redis, err := database.NewRedis(ctx, cfg.Redis.Addr(), cfg.Redis.Password)
	if err != nil {
		postgres.Close()
		return nil, err
	}

	// repositories
	linkRepo := link.NewPostgresRepository(postgres)

	// caches
	linkCache := link.NewRedisCache(redis, 24*time.Hour)

	// services
	apikeyService := apikey.NewService(cfg.App.APIKey)
	healthService := health.NewService(postgres, redis)
	linkService := link.NewService(
		link.ServiceConfig{
			BaseURL:         cfg.App.BaseURL,
			ShortCodeLength: cfg.App.ShortCodeLength,
		},
		linkRepo,
		linkCache,
	)

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

	a.closeOnce.Do(func() {
		if a.postgres != nil {
			a.postgres.Close()
		}

		if a.redis != nil {
			if err := a.redis.Close(); err != nil {
				errs = append(errs, err)
			}
		}
	})

	return errors.Join(errs...)
}
