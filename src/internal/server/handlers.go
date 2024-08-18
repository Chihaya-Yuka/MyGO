package server

import (
	"net/http"
)

// handlePut handles the PUT request.
func handlePut(w http.ResponseWriter, r *http.Request) {
	// Handle the PUT request.
	key := r.URL.Query().Get("key")
	value := r.URL.Query().Get("value")

	if err := s.store.Put(key, value); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// handleGet handles the GET request.
func handleGet(w http.ResponseWriter, r *http.Request) {
	// Handle the GET request.
	key := r.URL.Query().Get("key")

	value, err := s.store.Get(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Write([]byte(value))
}

// handleDelete handles the DELETE request.
func handleDelete(w http.ResponseWriter, r *http.Request) {
	// Handle the DELETE request.
	key := r.URL.Query().Get("key")

	if err := s.store.Delete(key); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
