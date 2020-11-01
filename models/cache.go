package models

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var Cache = cache.New(1*time.Minute, 2*time.Minute)

type CacheEntry struct {
	Key    string  `json:"-"`
	Action string  `json:"action"`
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	Answer float64 `json:"answer"`
	Cached bool    `json:"cached"`
}

func CacheSet(key string, entry interface{}) bool {
	Cache.Set(key, entry, cache.DefaultExpiration)
	return true
}

func CacheGet(key string) (CacheEntry, bool) {
	var entry CacheEntry
	var found bool
	data, found := Cache.Get(key)
	if found {
		entry = data.(CacheEntry)
	}
	return entry, found
}
