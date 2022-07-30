package storage

import (
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
)

var (
	redisCacheOnce  sync.Once
	redisClientOnce sync.Once
	redisCache      *cache.Cache
	redisClient     *redis.Client
)

func GetRedisClient(address string) *redis.Client {
	if redisClient == nil {
		redisClientOnce.Do(func() {
			fmt.Println("[storage][redis]--> Creating single redis client...")
			redisClient = redis.NewClient(&redis.Options{
				Addr: address,
			})
		})
	} else {
		fmt.Println("[storage][redis]--> redis client already created.")
	}
	return redisClient
}

func GetRedisCache(address string) *cache.Cache {
	if redisCache == nil {
		redisCacheOnce.Do(func() {
			fmt.Println("[storage][redis]--> Creating single redis cache...")
			redisCache = cache.New(&cache.Options{
				Redis:      GetRedisClient(address),
				LocalCache: cache.NewTinyLFU(1000, time.Minute),
			})
		})
	} else {
		fmt.Println("[storage][redis]--> redis cache already created.")
	}
	return redisCache
}
