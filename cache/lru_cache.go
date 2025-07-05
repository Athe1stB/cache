package cache

import (
	"athe1stb/cache/store"
	"container/list"
)

type LruCache[K comparable, V any] struct {
	capacity   int
	cacheStore *store.CacheStore[K, V]
	keyOrder   *list.List
	elementMap map[K]*list.Element
}

func NewLruCache[K comparable, V any](capacity int, cs store.CacheStore[K, V]) *LruCache[K, V] {
	return &LruCache[K, V]{
		capacity:   capacity,
		cacheStore: &cs,
		keyOrder:   list.New(),
		elementMap: make(map[K]*list.Element),
	}
}

func (c *LruCache[K, V]) Get(key K) (V, error) {
	val, err := (*c.cacheStore).GetValue(key)
	if err == nil {
		if curElement, present := c.elementMap[key]; present {
			c.keyOrder.MoveToBack(curElement)
		}
	}
	return val, err
}

func (c *LruCache[K, V]) Put(key K, value V) {
	if (*c.cacheStore).IsPresent(key) {
		c.Get(key)
		(*c.cacheStore).PutValue(key, value)
	} else {
		if c.capacity == (*c.cacheStore).Size() {
			c.Evict()
		}
		el := c.keyOrder.PushBack(key)
		c.elementMap[key] = el
		(*c.cacheStore).PutValue(key, value)
	}
}

func (c *LruCache[K, V]) Evict() {
	lruKey := c.keyOrder.Front()
	c.keyOrder.Remove(lruKey)
	key := (lruKey.Value).(K)
	(*c.cacheStore).Remove(key)
	delete(c.elementMap, key)
}
