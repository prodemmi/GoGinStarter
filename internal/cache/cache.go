package cache

import (
	"GoGinStarter/internal/config"
	"GoGinStarter/internal/log"
	"time"
)

type Cache interface{}

type cache struct {
	driver Driver
}

func (c *cache) Set(key string, value interface{}, ttl time.Duration) error {
	return c.driver.Set(key, value, ttl)
}

func (c *cache) Get(key string) (interface{}, error) {
	return c.driver.Get(key)
}

func (c *cache) Delete(key string) error {
	return c.driver.Delete(key)
}

func ProvideCache(config *config.Config, log log.Log) Cache {
	var driver Driver
	switch config.Cache.Driver {
	case "file":
		driver = NewFileCacheDriver("storage/cache")
	case "redis":
		d, err := NewRedisCacheDriver(config.Cache.Host, config.Cache.Password, config.Cache.Database)
		if err != nil {
			log.Error(err.Error())
			return nil
		}
		driver = d
	default:
		log.Error("Unknown cache driver " + config.Cache.Driver)
		return nil
	}

	return &cache{driver}
}
