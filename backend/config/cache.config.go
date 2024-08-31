package config

import (
	"fmt"
	"time"

	"github.com/jellydator/ttlcache/v2"
)

const (
	CacheNotFound = ttlcache.ErrNotFound
	CacheExpired  = ttlcache.ErrClosed
)

var cache *ttlcache.Cache

func InitCache() {
	cache = ttlcache.NewCache()
	fmt.Println("cache object: ", cache)
	cache.SetExpirationCallback(func(key string, value interface{}) {
		// if key expired, then cache will be removed
		defer recover()
		if err := cache.Remove(key); err == CacheNotFound {
			fmt.Printf("cache [%s] not found", key)
			return
		}
	})
}

// this method return initialized cache instance
func GetCacheTTL() *ttlcache.Cache {
	return cache
}

func SetKey(key string, value any, dura time.Duration) error {
	return cache.SetWithTTL(key, value, dura)
}

func GetKeyValue(key string) (interface{}, error) {
	return cache.Get(key)
}

func RMCache(key string) error {
	return cache.Remove(key)
}
