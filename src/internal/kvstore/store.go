package kvstore

import (
	"sync"
)

// Store represents a key-value store.
type Store struct {
	mu   sync.RWMutex
	data map[string][]byte
}

// NewStore returns a new instance of Store.
func NewStore() *Store {
	return &Store{
		data: make(map[string][]byte),
	}
}

// Get returns the value associated with the given key.
func (s *Store) Get(key string) ([]byte, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	value, ok := s.data[key]
	return value, ok
}

// Set sets the value associated with the given key.
func (s *Store) Set(key string, value []byte) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = value
}

// Delete deletes the value associated with the given key.
func (s *Store) Delete(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.data, key)
}

// Keys returns all keys in the store.
func (s *Store) Keys() []string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	keys := make([]string, 0, len(s.data))
	for key := range s.data {
		keys = append(keys, key)
	}
	return keys
}
