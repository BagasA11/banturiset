package config

import (
	"github.com/jellydator/ttlcache/v2"
)

const CacheNotFound = ttlcache.ErrNotFound

var cache *ttlcache.Cache

func InitCache() {
	cache = ttlcache.NewCache()
}

func GetCacheTTL() *ttlcache.Cache {
	return cache
}
