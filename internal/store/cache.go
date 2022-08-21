package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisStore struct {
	*redis.Client
}

func InitCache() *RedisStore {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return &RedisStore{rdb}
}

func (c *RedisStore) Cache(key string, value interface{}, duration time.Duration) error {
	err := c.Set(context.Background(), key, value, duration).Err()

	return err
}

func (c *RedisStore) ShortCache(key string, value interface{}) error {
	err := c.Set(context.Background(), key, value, 0).Err()

	return err
}

func (c *RedisStore) LongCache(key string, value interface{}) error {
	err := c.Set(context.Background(), key, value, 0).Err()

	return err
}
