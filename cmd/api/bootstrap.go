package main

import (
	"context"
	"log"

	"github.com/amirghafdurzadeh/golink/internal/config"
	"github.com/amirghafdurzadeh/golink/internal/database"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

func mustPostgres(ctx context.Context, cfg *config.Config) *pgxpool.Pool {
	pool, err := database.NewPostgres(ctx, cfg.PostgresURL)
	if err != nil {
		log.Fatal(err)
	}

	return pool
}

func mustRedis(ctx context.Context, cfg *config.Config) *redis.Client {
	client, err := database.NewRedis(ctx, cfg.RedisAddr, cfg.RedisPassword)
	if err != nil {
		log.Fatal(err)
	}

	return client
}
