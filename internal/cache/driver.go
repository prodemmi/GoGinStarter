package cache

import "time"

type Driver interface {
	Set(key string, value interface{}, ttl time.Duration) error
	Get(key string) (interface{}, error)
	Delete(key string) error
}
