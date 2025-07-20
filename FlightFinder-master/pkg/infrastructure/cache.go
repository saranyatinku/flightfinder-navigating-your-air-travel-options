package infrastructure

import "github.com/gin-gonic/gin"

// CacheClient is for find connection requests caching
type CacheClient interface {
	CachePage(handler gin.HandlerFunc) gin.HandlerFunc
}

// NullCacheClient provides a no-op implementation of CacheClient
type NullCacheClient struct {
}

func (c *NullCacheClient) CachePage(handler gin.HandlerFunc) gin.HandlerFunc {
	return handler // no caching
}
