package cache

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisStore struct {
	*redis.Client
}

func InitCache() *RedisStore {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_CONNECT"),
		Password: os.Getenv("REDIS_PASS"), // no password set
		DB:       0,                       // use default DB
	})

	return &RedisStore{rdb}
}

func (c *RedisStore) Cache(key string, value interface{}, duration int) error {
	var ctx = context.Background()

	// Set time in second
	dur := time.Duration(duration) * time.Second

	err := c.Set(ctx, key, value, dur).Err()

	return err
}

func (c *RedisStore) ShortCache(key string, value interface{}) error {
	var ctx = context.Background()

	// Set time in second
	intVar, err := strconv.Atoi(os.Getenv("REDIS_SHORT_CACHE"))
	if err != nil {
		return err
	}
	dur := time.Duration(intVar) * time.Second
	err = c.Set(ctx, key, value, dur).Err()

	return err
}

func (c *RedisStore) LongCache(key string, value interface{}) error {
	var ctx = context.Background()

	// Set time in second
	intVar, err := strconv.Atoi(os.Getenv("REDIS_LONG_CACHE"))
	if err != nil {
		return err
	}

	dur := time.Duration(intVar) * time.Second
	err = c.Set(ctx, key, value, dur).Err()

	return err
}
