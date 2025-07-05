package store

type CacheStore[K any, V any] interface {
	GetValue(key K) (V, error)
	PutValue(key K, value V)
	IsPresent(key K) bool
	Remove(key K)
	Size() int
}
