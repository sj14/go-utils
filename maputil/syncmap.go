package maputil

import "sync"

// based on https://www.reddit.com/r/golang/comments/twucb0/comment/j4x7xbx
type SyncMap[K comparable, V any] struct {
	innerMap sync.Map
}

func Make[K comparable, V any]() SyncMap[K, V] {
	return SyncMap[K, V]{}
}

func (m *SyncMap[K, V]) Store(key K, value V) {
	m.innerMap.Store(key, value)
}

func (m *SyncMap[K, V]) Contains(key K) bool {
	_, ok := m.innerMap.Load(key)
	return ok
}

func (m *SyncMap[K, V]) Delete(key K) {
	m.innerMap.Delete(key)
}

func (m *SyncMap[K, V]) Load(key K) (value V, ok bool) {
	v, ok := m.innerMap.Load(key)
	if !ok {
		return value, ok
	}
	return v.(V), ok
}

func (m *SyncMap[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	v, loaded := m.innerMap.LoadAndDelete(key)
	if !loaded {
		return value, loaded
	}
	return v.(V), loaded
}

func (m *SyncMap[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	a, loaded := m.innerMap.LoadOrStore(key, value)
	return a.(V), loaded
}

func (m *SyncMap[K, V]) Range(f func(key K, value V) bool) {
	m.innerMap.Range(func(key, value any) bool {
		return f(key.(K), value.(V))
	})
}
