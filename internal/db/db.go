package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type Database struct {
	Pool  *pgxpool.Pool
	Redis *redis.Client
}

func New(ctx context.Context, postgresURL, redisAddr, redisPassword string) (*Database, error) {
	pool, err := pgxpool.New(ctx, postgresURL)
	if err != nil {
		return nil, fmt.Errorf("failed to create postgres pool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping postgres: %w", err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       0,
	})

	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to ping redis: %w", err)
	}

	log.Println("Connected to Postgres and Redis")

	return &Database{
		Pool:  pool,
		Redis: rdb,
	}, nil
}

func (db *Database) Close() {
	db.Pool.Close()
	db.Redis.Close()
}
