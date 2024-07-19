package gomemcache

import (
	"fmt"
	"sync"
	"time"
)

type CacheItem struct {
	Value 		interface{}
	addedTime 	time.Time
	duration 	time.Duration
}

type Cache interface {
	Set(key string, value CacheItem)
	Get(key string)
	Remove(key string)
	Clear()
}

type MemCache struct {
	options map[string]CacheItem
	cleanupTime time.Duration
	mu sync.Mutex
}

func NewMemCache() MemCache {
	return MemCache {
		options: make(map[string]CacheItem),
	}
}

func (m *MemCache) Set(key string, value interface{}, duration time.Duration) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.options[key] = CacheItem{
		Value: value,
		addedTime: time.Now(),
		duration: duration,
	}

	fmt.Printf("Added: %s\n", key)
}

func (m *MemCache) Get(key string) (interface{}, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	item, found := m.options[key]
	if found {
		return item.Value, found
	}

	return nil, false
}

func (m *MemCache) Remove(key string) {
	delete(m.options, key)
}

func (m *MemCache) CleanAll() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.options = make(map[string]CacheItem)
}

func (m *MemCache) CheckDuration() {
	now := time.Now()

	for key, value := range m.options {
		if now.After(value.addedTime.Add(value.duration)) {
			delete(m.options, key)
		}
	}
}

func (m *MemCache) StartCleanup(interval time.Duration) {
	m.cleanupTime = interval
	ticker := time.NewTicker(interval)

	go func() {
		for range ticker.C {
			m.CleanAll()
			fmt.Println("Cache cleaned")
		}
	}()
}

