package cache

type Cache[K any, V any] interface {
	Get(K) (K, error)
	Put(K, V)
	Evict()
}
