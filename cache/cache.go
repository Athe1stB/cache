package cache

type Cache[K any, V any] interface {
	Get(K) (K, error)
	Put(K, V)
	Evict()
}

type TtlCache[K comparable, V any] interface {
	Cache[K, V]
	RemoveExpiredEntries()
}
