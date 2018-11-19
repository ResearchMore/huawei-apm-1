package apmcache

import cache "github.com/patrickmn/go-cache"

const (
	//DefaultExpireTime default expiry time is kept as 0
	DefaultExpireTime = 0
	// CleanupInterval default clean up time is kept as 0
	CleanupInterval = 0
)

// KPIWorkerCache  key: use src name,dest name and transactionType
var KPIWorkerCache *cache.Cache

func init() {
	KPIWorkerCache = initCache()
	KPIWorkerCache.DeleteExpired()
}

// NewCache return a new cache
func NewCache() *cache.Cache {
	return initCache()
}
func initCache() *cache.Cache { return cache.New(DefaultExpireTime, CleanupInterval) }
