package cacher

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	memory map[string]interface{}
	mu     sync.RWMutex
}

func New() Cache {
	return Cache{
		memory: make(map[string]interface{}),
	}
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	c.mu.Lock()
	c.memory[key] = value
	c.mu.Unlock()

	go func(key string, ttl time.Duration) {
		time.Sleep(ttl)
		c.mu.Lock()
		defer c.mu.Unlock()
		if _, ok := c.memory[key]; ok {
			c.delete(key)
		}
	}(key, ttl)
}

func (c *Cache) Get(key string) (interface{}, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	val, ok := c.memory[key]
	if !ok {
		return nil, fmt.Errorf("not found")
	}

	return val, nil
}

func (c *Cache) delete(key string) {
	delete(c.memory, key)
}
