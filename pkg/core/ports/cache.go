package ports

type CacheRepository[K comparable, T any] interface {
	GetCached(key K) (T, bool)
	SetCached(key K, value T)
	ClearCache()
}

type CacheClient[K comparable, T any] interface {
	Get(key K) (T, bool)
	Set(key K, value T)
	Clear()
}
