package link

import (
	"context"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
)

var errCacheMiss = errors.New("cache miss")

type redisCache struct {
	client *redis.Client
	ttl    time.Duration
}

func NewRedisCache(client *redis.Client, ttl time.Duration) Cache {
	return &redisCache{
		client: client,
		ttl:    ttl,
	}
}

func (c *redisCache) Get(ctx context.Context, code string) (string, error) {
	targetURL, err := c.client.Get(ctx, code).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", errCacheMiss
		}

		return "", err
	}

	return targetURL, nil
}

func (c *redisCache) Set(ctx context.Context, code string, targetURL string) error {
	return c.client.Set(ctx, code, targetURL, c.ttl).Err()
}

func (c *redisCache) Delete(ctx context.Context, code string) error {
	return c.client.Del(ctx, code).Err()
}
