package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// Config represents the configuration of the K-V database.
type Config struct {
	Addr          string `json:"addr"`
	LogLevel      string `json:"log_level"`
	StoreCapacity int    `json:"store_capacity"`
	StoreEngine   string `json:"store_engine"`
	StorePath     string `json:"store_path"`
}

// DefaultConfig returns the default configuration.
func DefaultConfig() *Config {
	return &Config{
		Addr:          ":8080",
		LogLevel:      "info",
		StoreCapacity: 1000000,
		StoreEngine:   "memory",
		StorePath:     "./store",
	}
}

// LoadConfig loads the configuration from a file.
func LoadConfig(filename string) (*Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

// SaveConfig saves the configuration to a file.
func SaveConfig(filename string, config *Config) error {
	data, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, data, 0644)
}

// ValidateConfig validates the configuration.
func ValidateConfig(config *Config) error {
	if config.Addr == "" {
		return fmt.Errorf("addr is required")
	}

	if config.LogLevel != "debug" && config.LogLevel != "info" && config.LogLevel != "warn" && config.LogLevel != "error" {
		return fmt.Errorf("invalid log level: %s", config.LogLevel)
	}

	if config.StoreCapacity <= 0 {
		return fmt.Errorf("store capacity must be greater than 0")
	}

	if config.StoreEngine != "memory" && config.StoreEngine != "disk" {
		return fmt.Errorf("invalid store engine: %s", config.StoreEngine)
	}

	if config.StorePath == "" {
		return fmt.Errorf("store path is required")
	}

	return nil
}
