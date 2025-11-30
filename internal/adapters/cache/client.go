package cache

import (
	"container/list"
	"sync"
	"time"

	"github.com/Bikes2Road/bikes-compass/internal/core/ports"
)

type CacheEntry[T any] struct {
	value     T
	timestamp time.Time
	element   *list.Element
}

type CacheClient[K comparable, T any] interface {
	Get(key K) (T, bool)
	Set(key K, value T)
	Clear()
}

type LRUCache[K comparable, T any] struct {
	entries  map[K]CacheEntry[T]
	ttl      time.Duration
	capacity int
	lruList  *list.List
	mutext   sync.RWMutex
}

func NewCacheClient(capacity int, ttl time.Duration) ports.CacheClient[string, any] {
	// Puedes ajustar capacidad y TTL a tus necesidades
	return NewLRUCache[string, any](capacity, ttl*time.Minute)
}

func NewLRUCache[K comparable, T any](capacity int, ttl time.Duration) *LRUCache[K, T] {
	return &LRUCache[K, T]{
		entries:  make(map[K]CacheEntry[T]),
		ttl:      ttl,
		capacity: capacity,
		lruList:  list.New(),
		mutext:   sync.RWMutex{},
	}
}

func (c *LRUCache[K, T]) Get(key K) (T, bool) {
	c.mutext.Lock()
	defer c.mutext.Unlock()

	entry, exists := c.entries[key]
	if !exists {
		var zero T
		return zero, false
	}

	if time.Since(entry.timestamp) > c.ttl {
		c.lruList.Remove(entry.element)
		delete(c.entries, key)
		var zero T
		return zero, false
	}

	c.lruList.MoveToFront(entry.element)

	return entry.value, true
}

func (c *LRUCache[K, T]) Set(key K, value T) {
	c.mutext.Lock()
	defer c.mutext.Unlock()

	if entry, exists := c.entries[key]; exists {
		entry.value = value
		entry.timestamp = time.Now()
		c.entries[key] = entry
		c.lruList.MoveToFront(entry.element)
		return
	}

	if c.lruList.Len() >= c.capacity {
		oldest := c.lruList.Back()
		c.lruList.Remove(oldest)
		delete(c.entries, oldest.Value.(K))
	}

	c.entries[key] = CacheEntry[T]{value: value, timestamp: time.Now(), element: c.lruList.PushFront(key)}
}

func (c *LRUCache[K, T]) Clear() {
	c.mutext.Lock()
	defer c.mutext.Unlock()
	c.entries = make(map[K]CacheEntry[T])
	c.lruList.Init()
}
