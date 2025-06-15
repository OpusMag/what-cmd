package discovery

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"what-cmd/internal/config"
	"what-cmd/internal/models"
)

type Cache struct {
	config    config.CacheConfig
	cacheFile string
}

type CacheData struct {
	Items     []models.Item `json:"items"`
	Timestamp time.Time     `json:"timestamp"`
	Version   string        `json:"version"`
}

func NewCache(config config.CacheConfig) (*Cache, error) {
	if !config.Enabled {
		return &Cache{config: config}, nil
	}

	if err := os.MkdirAll(config.Directory, 0755); err != nil {
		return nil, fmt.Errorf("failed to create cache directory: %w", err)
	}

	cacheFile := filepath.Join(config.Directory, "discovered_items.json")

	return &Cache{
		config:    config,
		cacheFile: cacheFile,
	}, nil
}

func (c *Cache) StoreItems(items []models.Item) error {
	if !c.config.Enabled {
		return nil
	}

	data := CacheData{
		Items:     items,
		Timestamp: time.Now(),
		Version:   "1.0",
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal cache data: %w", err)
	}

	if err := os.WriteFile(c.cacheFile, jsonData, 0644); err != nil {
		return fmt.Errorf("failed to write cache file: %w", err)
	}

	return nil
}

func (c *Cache) GetItems() ([]models.Item, error) {
	if !c.config.Enabled {
		return nil, fmt.Errorf("cache is disabled")
	}

	if _, err := os.Stat(c.cacheFile); os.IsNotExist(err) {
		return nil, fmt.Errorf("cache file does not exist")
	}

	data, err := os.ReadFile(c.cacheFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read cache file: %w", err)
	}

	var cacheData CacheData
	if err := json.Unmarshal(data, &cacheData); err != nil {
		return nil, fmt.Errorf("failed to unmarshal cache data: %w", err)
	}

	ttl := time.Duration(c.config.TTLHours) * time.Hour
	if time.Since(cacheData.Timestamp) > ttl {
		return nil, fmt.Errorf("cache has expired")
	}

	return cacheData.Items, nil
}

func (c *Cache) ClearCache() error {
	if !c.config.Enabled {
		return nil
	}

	if err := os.Remove(c.cacheFile); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to remove cache file: %w", err)
	}

	return nil
}
