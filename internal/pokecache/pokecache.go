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

func NewCache(time.Duration) Cache {
	newCache := Cache{
		mutex: &sync.Mutex{},
	}
	newCache.reapLoop()
	return newCache

}

func (c Cache) Add(key string val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.cached[key] = val
}
func (c Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if val, ok := c.cached[key]; !ok{
		return []byte{}, false
	}
	return val, true
}
func (c Cache) reapLoop(interval time.Duration) {
	ticker := time.newTicker(interval)
	for {
		select {
		case t := <-ticker.C:
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