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

func (c *Cache) StoreItems(items []models.Item) error {
	if !c.config.Enabled {
		return nil
	}

	data := CacheData{
		Items:     items,
		Timestamp: time.Now(),
		Version:   "1.1",
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

func (c *Cache) GetCacheInfo() (map[string]interface{}, error) {
	if !c.config.Enabled {
		return map[string]interface{}{
			"enabled": false,
		}, nil
	}

	info := map[string]interface{}{
		"enabled":   true,
		"file_path": c.cacheFile,
		"exists":    false,
	}

	if stat, err := os.Stat(c.cacheFile); err == nil {
		info["exists"] = true
		info["size_bytes"] = stat.Size()
		info["modified"] = stat.ModTime()

		if data, err := os.ReadFile(c.cacheFile); err == nil {
			var cacheData CacheData
			if err := json.Unmarshal(data, &cacheData); err == nil {
				info["item_count"] = len(cacheData.Items)
				info["timestamp"] = cacheData.Timestamp
				info["version"] = cacheData.Version

				ttl := time.Duration(c.config.TTLHours) * time.Hour
				info["expires_at"] = cacheData.Timestamp.Add(ttl)
				info["is_expired"] = time.Since(cacheData.Timestamp) > ttl
				info["age_hours"] = time.Since(cacheData.Timestamp).Hours()
			}
		}
	}

	return info, nil
}
