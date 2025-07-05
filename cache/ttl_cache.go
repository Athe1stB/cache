package cache

type TtlCache[K comparable, V any] interface {
	Cache[K, V]
	removeExpired()
}
