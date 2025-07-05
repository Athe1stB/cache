package store

import "fmt"

type InMemoryStore[K comparable, V any] struct {
	cache map[K]V
}

func NewInMemoryStore[K comparable, V any]() *InMemoryStore[K, V] {
	return &InMemoryStore[K, V]{
		cache: make(map[K]V),
	}
}

func (c *InMemoryStore[K, V]) GetValue(key K) (V, error) {
	val := c.cache[key]
	if _, ok := c.cache[key]; ok {
		return val, nil
	} else {
		return val, fmt.Errorf("value not present in cache")
	}
}

func (c *InMemoryStore[K, V]) PutValue(key K, val V) {
	c.cache[key] = val
}

func (c *InMemoryStore[K, V]) IsPresent(key K) bool {
	_, ok := c.cache[key]
	return ok
}

func (c *InMemoryStore[K, V]) Size() int {
	return len(c.cache)
}

func (c *InMemoryStore[K, V]) Remove(key K) {
	delete(c.cache, key)
}
