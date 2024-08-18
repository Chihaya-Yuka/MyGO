package server

import (
	"net/http"
)

// APIHandler represents an API handler.
type APIHandler func(http.ResponseWriter, *http.Request)

// apiHandlers maps API paths to their handlers.
var apiHandlers = map[string]APIHandler{
	"/api/put":    handlePut,
	"/api/get":    handleGet,
	"/api/delete": handleDelete,
}

// registerAPIHandlers registers API handlers.
func (s *Server) registerAPIHandlers() {
	for path, handler := range apiHandlers {
		http.HandleFunc(path, handler)
	}
}
