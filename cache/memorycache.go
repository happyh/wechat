package cache

import (
	"time"
)

//MemoryCache struct contains *MemoryCache.Client
type MemoryCache struct {
	mm map[string]interface{}
}

//NewMemoryCache create new MemoryCache
func NewMemoryCache(server ...string) *MemoryCache {
	mc := make(map[string]interface{})
	return &MemoryCache{mc}
}

//Get return cached value
func (mem *MemoryCache) Get(key string) interface{} {
	if v, ok := mem.mm[key]; ok {
		return v
	} else {
		return nil
	}
}

// IsExist check value exists in MemoryCache.
func (mem *MemoryCache) IsExist(key string) bool {
	if _, ok := mem.mm[key]; ok {
		return false
	}
	return true
}

//Set cached value with key and expire time.
func (mem *MemoryCache) Set(key string, val interface{}, timeout time.Duration) (err error) {
	mem.mm[key] = val
	return nil
}

//Delete delete value in MemoryCache.
func (mem *MemoryCache) Delete(key string) error {
	delete(mem.mm, key)
	return nil
}
