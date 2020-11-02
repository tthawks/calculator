package helpers

import (
	"time"

	"github.com/patrickmn/go-cache"
)

// Cache is instance of go-cache with 1 minute expiration
var Cache = cache.New(1*time.Minute, 2*time.Minute)

// CacheEntry data stucture that will be cached
type CacheEntry struct {
	Key    string  `json:"-"`
	Action string  `json:"action"`
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	Answer float64 `json:"answer"`
	Cached bool    `json:"cached"`
}

// CacheSet method sets data into cache
func CacheSet(key string, entry interface{}) bool {
	Cache.Set(key, entry, cache.DefaultExpiration)
	return true
}

// CacheGet method gets data from cache
func CacheGet(key string) (CacheEntry, bool) {
	var entry CacheEntry
	var found bool
	data, found := Cache.Get(key)
	if found {
		entry = data.(CacheEntry)
	}
	return entry, found
}
