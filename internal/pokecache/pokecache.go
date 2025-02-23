package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cached map[string]cachedEntry
	mutex  *sync.Mutex
}
type cachedEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	newCache := Cache{
		mutex:  &sync.Mutex{},
		cached: map[string]cachedEntry{},
	}
	go newCache.reapLoop(interval)
	return newCache

}

func (c Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.cached[key] = cachedEntry{time.Now(), val}
}
func (c Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	entry, ok := c.cached[key]
	if !ok {
		return []byte{}, false
	}
	return entry.val, true
}
func (c Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for {
		for range ticker.C {
			c.mutex.Lock()
			for key, val := range c.cached {
				if val.createdAt.Before(time.Now().Add(-interval)) {
					delete(c.cached, key)
				}
			}
			c.mutex.Unlock()
		}
	}
}
