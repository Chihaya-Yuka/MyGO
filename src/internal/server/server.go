package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/Chihaya-Yuka/mygo/internal/config"
	"github.com/Chihaya-Yuka/mygo/internal/kvstore"
	"github.com/Chihaya-Yuka/mygo/internal/kvstore/memory"
	"github.com/Chihaya-Yuka/mygo/pkg/logger"
)

// Server represents the K-V database server.
type Server struct {
	config  *config.Config
	store   kvstore.Store
	logger  *logger.Logger
	httpSrv *http.Server
}

// NewServer returns a new instance of the server.
func NewServer(cfg *config.Config) (*Server, error) {
	// Initialize the logger.
	logger, err := logger.NewLogger(cfg.LogLevel)
	if err != nil {
		return nil, err
	}

	// Initialize the K-V store.
	store, err := memory.NewMemoryStore(cfg.StoreCapacity)
	if err != nil {
		return nil, err
	}

	// Initialize the HTTP server.
	httpSrv := &http.Server{
		Addr: cfg.Addr,
	}

	return &Server{
		config:  cfg,
		store:   store,
		logger:  logger,
		httpSrv: httpSrv,
	}, nil
}

// Start starts the server.
func (s *Server) Start() error {
	// Start the HTTP server.
	s.logger.Info("Starting server on", s.config.Addr)
	go func() {
		if err := s.httpSrv.ListenAndServe(); err != nil {
			s.logger.Error("Failed to start server:", err)
		}
	}()

	// Register API handlers.
	s.registerAPIHandlers()

	return nil
}

// Stop stops the server.
func (s *Server) Stop() error {
	// Stop the HTTP server.
	s.logger.Info("Stopping server")
	return s.httpSrv.Close()
}

// registerAPIHandlers registers API handlers.
func (s *Server) registerAPIHandlers() {
	// Register API handlers.
	http.HandleFunc("/api/put", s.handlePut)
	http.HandleFunc("/api/get", s.handleGet)
	http.HandleFunc("/api/delete", s.handleDelete)
}

// handlePut handles the PUT request.
func (s *Server) handlePut(w http.ResponseWriter, r *http.Request) {
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
func (s *Server) handleGet(w http.ResponseWriter, r *http.Request) {
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
func (s *Server) handleDelete(w http.ResponseWriter, r *http.Request) {
	// Handle the DELETE request.
	key := r.URL.Query().Get("key")

	if err := s.store.Delete(key); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
