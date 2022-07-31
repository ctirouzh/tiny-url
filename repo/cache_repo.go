package repo

import (
	"context"

	"github.com/ctirouzh/tiny-url/config"
	"github.com/ctirouzh/tiny-url/model"
	"github.com/go-redis/cache/v8"
)

type CacheRepository struct {
	cache  *cache.Cache
	config config.Redis
}

func NewCacheRepository(cache *cache.Cache, config config.Redis) *CacheRepository {
	return &CacheRepository{cache: cache, config: config}
}

func (cr *CacheRepository) SetURL(url *model.URL) {
	err := cr.cache.Set(&cache.Item{
		Ctx:   context.TODO(),
		Key:   url.Hash,
		Value: url,
		TTL:   cr.config.TTL,
	})
	if err != nil {
		return
	}
}

func (cr *CacheRepository) GetURL(hash string) *model.URL {
	var url model.URL
	err := cr.cache.Get(context.TODO(), hash, &url)
	if err != nil {
		return nil
	}
	return &url
}
