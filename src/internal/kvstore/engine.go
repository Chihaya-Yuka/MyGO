package kvstore

import (
	"errors"
	"fmt"
	"log"
)

// Engine represents a database engine.
type Engine struct {
	store *Store
}

// NewEngine returns a new instance of Engine.
func NewEngine(store *Store) *Engine {
	return &Engine{
		store: store,
	}
}

// Start starts the database engine.
func (e *Engine) Start() {
	log.Println("Database engine started")
}

// Stop stops the database engine.
func (e *Engine) Stop() {
	log.Println("Database engine stopped")
}

// Get returns the value associated with the given key.
func (e *Engine) Get(key string) ([]byte, error) {
	value, ok := e.store.Get(key)
	if !ok {
		return nil, errors.New("key not found")
	}
	return value, nil
}

// Set sets the value associated with the given key.
func (e *Engine) Set(key string, value []byte) error {
	e.store.Set(key, value)
	return nil
}

// Delete deletes the value associated with the given key.
func (e *Engine) Delete(key string) error {
	e.store.Delete(key)
	return nil
}

// Keys returns all keys in the database.
func (e *Engine) Keys() ([]string, error) {
	keys := e.store.Keys()
	return keys, nil
}

// Stats returns statistics about the database.
func (e *Engine) Stats() string {
	return fmt.Sprintf("Number of keys: %d", len(e.store.Keys()))
}
