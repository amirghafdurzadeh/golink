package health

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

var (
	ErrPostgresUnhealthy = errors.New("postgres unhealthy")
	ErrRedisUnhealthy    = errors.New("redis unhealthy")
)

type Service interface {
	Check(ctx context.Context) error
}

type service struct {
	pgPool      *pgxpool.Pool
	redisClient *redis.Client
}

func NewService(pgPool *pgxpool.Pool, redisClient *redis.Client) Service {
	return &service{
		pgPool:      pgPool,
		redisClient: redisClient,
	}
}

func (s *service) Check(ctx context.Context) error {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	if err := s.pgPool.Ping(ctxWithTimeout); err != nil {
		return ErrPostgresUnhealthy
	}

	if err := s.redisClient.Ping(ctxWithTimeout).Err(); err != nil {
		return ErrRedisUnhealthy
	}

	return nil
}
