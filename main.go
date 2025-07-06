package main

import (
	"athe1stb/cache/cache"
	"athe1stb/cache/store"
	"fmt"
	"time"
)

func main() {
	var store store.CacheStore[int, int] = store.NewInMemoryStore[int, int]()
	cache := cache.NewLruCache(2, 1000, store)

	go cache.RemoveExpiredEntries()

	printValue(cache, 5)
	cache.Put(3, 5)
	cache.Put(1, 1)
	printValue(cache, 3)
	cache.Put(2, 2)
	time.Sleep(time.Second * 1)
	printValue(cache, 1)
	printValue(cache, 2)
	printValue(cache, 3)
	time.Sleep(time.Hour)
}

func printValue(cache cache.Cache[int, int], key int) {
	val, err := cache.Get(key)
	if err == nil {
		fmt.Println("Value: ", val)
	} else {
		fmt.Println(err)
	}

}
