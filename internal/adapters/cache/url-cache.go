package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type UrlCache struct {
	client *redis.Client
	ttl time.Duration
}

func NewUrlCache(client *redis.Client) *UrlCache {
	return &UrlCache{
		client: client,
		ttl: time.Hour * 24,
	}
}

func (c *UrlCache) Set(ctx context.Context, code string, originalUrl string) error {
    return c.client.Set(ctx, code, originalUrl, c.ttl).Err()
}

func (c *UrlCache) Get(ctx context.Context, code string) (string, error) {
    return c.client.Get(ctx, code).Result()
}

func (c *UrlCache) Delete(ctx context.Context, code string) error {
    return c.client.Del(ctx, code).Err()
}