package cache

import "github.com/Bikes2Road/bikes-compass/internal/core/ports"

type CacheRepository struct {
	client ports.CacheClient[string, any]
}

func NewCacheRepository(client ports.CacheClient[string, any]) ports.CacheRepository[string, any] {
	return &CacheRepository{client: client}
}

func (r *CacheRepository) GetCached(key string) (any, bool) {
	return r.client.Get(key)
}

func (r *CacheRepository) SetCached(key string, value any) {
	r.client.Set(key, value)
}

func (r *CacheRepository) ClearCache() {
	r.client.Clear()
}
