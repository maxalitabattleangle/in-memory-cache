package cacher

import "fmt"

type Cache struct {
	memory map[string]interface{}
}

func New() Cache {
	return Cache{
		memory: make(map[string]interface{}),
	}
}

func (c *Cache) Set(key string, value interface{}) {
	c.memory[key] = value
}

func (c Cache) Get(key string) interface{} {
	val, ok := c.find(key)

	if !ok {
		fmt.Printf("Элемент по ключу %s не найден.\n", key)
	}

	return val
}

func (c *Cache) Delete(key string) {
	if _, ok := c.find(key); ok {
		delete(c.memory, key)
	} else {
		fmt.Printf("Элемент по ключу %s не найден.\n", key)
	}
}

func (c *Cache) find(key string) (val interface{}, b bool) {
	val, ok := c.memory[key]

	if ok {
		b = true
	}

	return val, b
}
