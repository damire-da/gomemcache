package memcache

import (
	"time"
)

type CacheItem struct {
	Value interface{}
	Expiration int64
}

type MemCache struct {
	data map[string]CacheItem
	ttl time.Duration
}

func New(ttl time.Duration) *MemCache {
	return &MemCache{
		data: make(map[string]CacheItem),
		ttl: ttl,
	}
}

func (c *MemCache) Get(key string) (interface{}, bool) {
	item, found := c.data[key]
	
	if !found {
		return nil, false
	}
	
	if time.Now().Unix() > item.Expiration {
		delete(c.data, key)
		return nil, false
	}
	
	return item.Value, true
}

func (c *MemCache) Set(key string, value interface{}) {
	expiration := time.Now().Add(c.ttl).Unix()
	c.data[key] = CacheItem{
		Value:     value,
        Expiration: expiration,
	}
}
