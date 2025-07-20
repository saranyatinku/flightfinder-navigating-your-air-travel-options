package aws

import (
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
)

// RedisCacheClient implements CacheClient interface
type RedisCacheClient struct {
	store  persistence.CacheStore
	expire time.Duration
}

func NewCacheClient(addr, pass string, expire time.Duration) (*RedisCacheClient, error) {
	store := persistence.NewRedisCache(addr, pass, expire)

	// check if connection to Redis works
	err := store.Set("test-redis-connection", true, time.Duration(0))
	if err != nil {
		return nil, err
	}
	return &RedisCacheClient{store: store, expire: expire}, nil
}

func (c *RedisCacheClient) CachePage(handler gin.HandlerFunc) gin.HandlerFunc {
	return cache.CachePage(c.store, c.expire, handler)
}
