package memcache

type MemCache struct {
	data map[string]interface{}
}

func New() MemCache {
	return MemCache{
		data: make(map[string]interface{}),
	}
}

func (c MemCache) Get(key string) interface{} {
	return c.data[key]
}

func (c *MemCache) Set(key string, value interface{}) {
	c.data[key] = value
}

func (c *MemCache) Delete(key string) {
	delete(c.data, key)
}
