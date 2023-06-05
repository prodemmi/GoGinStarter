package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisCacheDriver struct {
	client *redis.Client
}

func NewRedisCacheDriver(host string, password string, database int) (*RedisCacheDriver, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       database,
	})
	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}
	return &RedisCacheDriver{client: client}, nil
}

func (c *RedisCacheDriver) Set(key string, value interface{}, ttl time.Duration) error {
	return c.client.Set(context.Background(), key, value, ttl).Err()
}

func (c *RedisCacheDriver) Get(key string) (interface{}, error) {
	return c.client.Get(context.Background(), key).Result()
}

func (c *RedisCacheDriver) Delete(key string) error {
	return c.client.Del(context.Background(), key).Err()
}
