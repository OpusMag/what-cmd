package discovery

import (
	"crypto/sha256"
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
	Items      []models.Item `json:"items"`
	Timestamp  time.Time     `json:"timestamp"`
	ConfigHash string        `json:"config_hash"`
	Version    string        `json:"version"`
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

func configHash(cfg config.DiscoveryConfig) string {
	data, _ := json.Marshal(cfg)
	sum := sha256.Sum256(data)
	return fmt.Sprintf("%x", sum)
}

func (c *Cache) StoreItems(items []models.Item, hash string) error {
	if !c.config.Enabled {
		return nil
	}

	data := CacheData{
		Items:      items,
		Timestamp:  time.Now(),
		ConfigHash: hash,
		Version:    "1.1",
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

func (c *Cache) GetItems(currentHash string) ([]models.Item, error) {
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

	if cacheData.ConfigHash != currentHash {
		return nil, fmt.Errorf("cache config hash mismatch - configuration has changed")
	}

	ttl := time.Duration(c.config.TTLHours) * time.Hour
	if time.Since(cacheData.Timestamp) > ttl {
		return nil, fmt.Errorf("cache expired (age: %v, TTL: %v)",
			time.Since(cacheData.Timestamp).Round(time.Hour), ttl)
	}

	if len(cacheData.Items) == 0 {
		return nil, fmt.Errorf("cache contains no items")
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

