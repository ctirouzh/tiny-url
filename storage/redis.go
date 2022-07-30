package storage

import (
	"fmt"
	"sync"

	"github.com/ctirouzh/tiny-url/config"
	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
)

var (
	redisCacheOnce  sync.Once
	redisClientOnce sync.Once
	redisCache      *cache.Cache
	redisClient     *redis.Client
)

func GetRedisClient(cfg config.Redis) *redis.Client {
	if redisClient == nil {
		redisClientOnce.Do(func() {
			fmt.Println("[storage][redis]--> Creating single redis client...")
			redisClient = redis.NewClient(&redis.Options{
				Addr: cfg.Address,
			})
		})
	} else {
		fmt.Println("[storage][redis]--> redis client already created.")
	}
	return redisClient
}

func GetRedisCache(cfg config.Redis) *cache.Cache {
	if redisCache == nil {
		redisCacheOnce.Do(func() {
			fmt.Println("[storage][redis]--> Creating single redis cache...")
			redisCache = cache.New(&cache.Options{
				Redis: GetRedisClient(cfg),
				// Cache "Size" keys for "TTL" minutes
				LocalCache: cache.NewTinyLFU(cfg.LFU.Size, cfg.LFU.TTL),
			})
		})
	} else {
		fmt.Println("[storage][redis]--> redis cache already created.")
	}
	return redisCache
}
