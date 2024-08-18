package config

import (
	"testing"
)

func TestDefaultConfig(t *testing.T) {
	config := DefaultConfig()
	if config.Addr != ":8080" {
		t.Errorf("expected addr to be :8080, but got %s", config.Addr)
	}
	if config.LogLevel != "info" {
		t.Errorf("expected log level to be info, but got %s", config.LogLevel)
	}
	if config.StoreCapacity != 1000000 {
		t.Errorf("expected store capacity to be 1000000, but got %d", config.StoreCapacity)
	}
	if config.StoreEngine != "memory" {
		t.Errorf("expected store engine to be memory, but got %s", config.StoreEngine)
	}
	if config.StorePath != "./store" {
		t.Errorf("expected store path to be ./store, but got %s", config.StorePath)
	}
}

func TestLoadConfig(t *testing.T) {
	filename := "test_config.json"
	config := &Config{
		Addr:          ":8081",
		LogLevel:      "debug",
		StoreCapacity: 500000,
		StoreEngine:   "disk",
		StorePath:     "./test_store",
	}
	if err := SaveConfig(filename, config); err != nil {
		t.Fatal(err)
	}
	defer os.Remove(filename)

	loadedConfig, err := LoadConfig(filename)
	if err != nil {
		t.Fatal(err)
	}
	if loadedConfig.Addr != config.Addr {
		t.Errorf("expected addr to be %s, but got %s", config.Addr, loadedConfig.Addr)
	}
	if loadedConfig.LogLevel != config.LogLevel {
		t.Errorf("expected log level to be %s, but got %s", config.LogLevel, loadedConfig.LogLevel)
	}
	if loadedConfig.StoreCapacity != config.StoreCapacity {
		t.Errorf("expected store capacity to be %d, but got %d", config.StoreCapacity, loadedConfig.StoreCapacity)
	}
	if loadedConfig.StoreEngine != config.StoreEngine {
		t.Errorf("expected store engine to be %s, but got %s", config.StoreEngine, loadedConfig.StoreEngine)
	}
	if loadedConfig.StorePath != config.StorePath {
		t.Errorf("expected store path to be %s, but got %s", config.StorePath, loadedConfig.StorePath)
	}
}

func TestValidateConfig(t *testing.T) {
	config := &Config{
		Addr:          "",
		LogLevel:      "invalid",
		StoreCapacity: 0,
		StoreEngine:   "invalid",
		StorePath:     "",
	}
	if err := ValidateConfig(config); err == nil {
		t.Errorf("expected validation to fail, but got no error")
	}
}
