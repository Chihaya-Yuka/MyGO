package kvstore

import (
	"fmt"
	"log"
	"runtime"
)

// MemoryStore represents a key-value store that uses memory as its backend.
type MemoryStore struct {
	*Store
}

// NewMemoryStore returns a new instance of MemoryStore.
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		Store: NewStore(),
	}
}

// Start starts the memory store.
func (m *MemoryStore) Start() {
	log.Println("Memory store started")
}

// Stop stops the memory store.
func (m *MemoryStore) Stop() {
	log.Println("Memory store stopped")
}

// Stats returns statistics about the memory store.
func (m *MemoryStore) Stats() string {
	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)
	return fmt.Sprintf("Memory usage: %d bytes", stats.Alloc)
}
