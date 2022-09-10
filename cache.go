package cache

import "sync"

// ICache is a common-cache interface.
type ICache[K comparable, V any] interface {
	Add(key K, value V) bool
	Get(key K) (value V, ok bool)
}

// Cache is a thread-safe cache.
type Cache[K comparable, V any] struct {
	cache ICache[K, V]

	lock sync.RWMutex
}

// New creates a new thread safe Cache.
func New[K comparable, V any](cache ICache[K, V]) *Cache[K, V] {
	return &Cache[K, V]{cache: cache}
}

// Add adds a value to the cache. Returns true if an eviction occurred.
func (c *Cache[K, V]) Add(key K, value V) bool {
	c.lock.Lock()
	defer c.lock.Unlock()

	return c.cache.Add(key, value)
}

// Get looks up a key's value from the cache.
func (c *Cache[K, V]) Get(key K) (value V, ok bool) {
	c.lock.Lock()
	defer c.lock.Unlock()

	return c.cache.Get(key)
}
