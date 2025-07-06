package cache

import (
	"athe1stb/cache/store"
	"container/list"
	"fmt"
	"sync/atomic"
	"time"
)

type LruEntry[K any] struct {
	insertionTime time.Time
	key           K
}

type LruCache[K comparable, V any] struct {
	capacity    int
	ttlInMillis int
	cacheStore  *store.CacheStore[K, V]
	keyOrder    *list.List
	elementMap  map[K]*list.Element
	ticker      *time.Ticker
}

func NewLruCache[K comparable, V any](capacity int, ttlInMillis int, cs store.CacheStore[K, V]) *LruCache[K, V] {
	return &LruCache[K, V]{
		capacity:    capacity,
		cacheStore:  &cs,
		ttlInMillis: ttlInMillis,
		keyOrder:    list.New(),
		elementMap:  make(map[K]*list.Element),
		ticker:      time.NewTicker(time.Second * 15),
	}
}

func (c *LruCache[K, V]) Get(key K) (V, error) {
	val, err := (*c.cacheStore).GetValue(key)
	if err == nil {
		if curElement, present := c.elementMap[key]; present {
			lruEntry := (curElement.Value).(LruEntry[K])
			if lruEntry.insertionTime.Add(time.Millisecond * time.Duration(c.ttlInMillis)).Before(time.Now()) {
				c.keyOrder.Remove(curElement)
				(*c.cacheStore).Remove(key)
				delete(c.elementMap, key)
				return val, fmt.Errorf("TTL expired for key %v", key)
			}
			c.keyOrder.MoveToBack(curElement)
		}
	}
	return val, err
}

func (c *LruCache[K, V]) Put(key K, value V) {
	newEntry := LruEntry[K]{key: key, insertionTime: time.Now()}
	if (*c.cacheStore).IsPresent(key) {
		c.keyOrder.Remove(c.elementMap[key])
	} else {
		if c.capacity == (*c.cacheStore).Size() {
			c.Evict()
		}
	}
	el := c.keyOrder.PushBack(newEntry)
	c.elementMap[key] = el
	(*c.cacheStore).PutValue(key, value)
}

func (c *LruCache[K, V]) Evict() {
	lruKey := c.keyOrder.Front()
	c.keyOrder.Remove(lruKey)
	lruEntry := (lruKey.Value).(LruEntry[K])
	(*c.cacheStore).Remove(lruEntry.key)
	delete(c.elementMap, lruEntry.key)
}

func (c *LruCache[K, V]) RemoveExpiredEntries() {
	for range c.ticker.C {
		fmt.Println("Starting to remove expired entries")
		expiredEntries := atomic.Int32{}
		for key, element := range c.elementMap {
			expiredEntries.Add(1)
			c.keyOrder.Remove(element)
			(*c.cacheStore).Remove(key)
			delete(c.elementMap, key)
		}
		fmt.Printf("Successfully removed %d expired entries\n", expiredEntries.Load())
	}
}
