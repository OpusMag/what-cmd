package discovery

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"what-cmd/internal/config"
	"what-cmd/internal/models"
)

// ─── Cache.StoreItems / GetItems ─────────────────────────────────────────────

func TestCache_StoreAndRetrieve(t *testing.T) {
	dir := t.TempDir()
	cfg := config.CacheConfig{
		Enabled:  true,
		Directory: dir,
		TTLHours: 24,
	}
	c, err := NewCache(cfg)
	if err != nil {
		t.Fatalf("NewCache: %v", err)
	}

	items := []models.Item{
		{Name: "git", Description: "version control"},
	}
	hash := "abc123"

	if err := c.StoreItems(items, hash); err != nil {
		t.Fatalf("StoreItems: %v", err)
	}

	got, err := c.GetItems(hash)
	if err != nil {
		t.Fatalf("GetItems: %v", err)
	}
	if len(got) != 1 || got[0].Name != "git" {
		t.Errorf("retrieved items mismatch: %v", got)
	}
}

func TestCache_ConfigHashInvalidation(t *testing.T) {
	dir := t.TempDir()
	cfg := config.CacheConfig{
		Enabled:  true,
		Directory: dir,
		TTLHours: 24,
	}
	c, err := NewCache(cfg)
	if err != nil {
		t.Fatalf("NewCache: %v", err)
	}

	items := []models.Item{{Name: "git", Description: "version control"}}
	if err := c.StoreItems(items, "abc"); err != nil {
		t.Fatalf("StoreItems: %v", err)
	}

	_, err = c.GetItems("xyz")
	if err == nil {
		t.Error("expected cache miss for mismatched config hash, got nil error")
	}
}

func TestCache_TTLExpiry(t *testing.T) {
	dir := t.TempDir()
	cfg := config.CacheConfig{
		Enabled:  true,
		Directory: dir,
		TTLHours: 1,
	}
	c, err := NewCache(cfg)
	if err != nil {
		t.Fatalf("NewCache: %v", err)
	}

	// Write a cache file with a timestamp 2 hours in the past (TTL is 1 hour).
	data := CacheData{
		Items:      []models.Item{{Name: "git", Description: "version control"}},
		Timestamp:  time.Now().Add(-2 * time.Hour),
		ConfigHash: "abc",
		Version:    "1.1",
	}
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		t.Fatalf("json.MarshalIndent: %v", err)
	}
	cacheFile := filepath.Join(dir, "discovered_items.json")
	if err := os.WriteFile(cacheFile, jsonData, 0644); err != nil {
		t.Fatalf("os.WriteFile: %v", err)
	}

	_, err = c.GetItems("abc")
	if err == nil {
		t.Error("expected cache miss for expired TTL, got nil error")
	}
}

func TestCache_Disabled(t *testing.T) {
	cfg := config.CacheConfig{
		Enabled: false,
	}
	c, err := NewCache(cfg)
	if err != nil {
		t.Fatalf("NewCache: %v", err)
	}

	_, err = c.GetItems("any-hash")
	if err == nil {
		t.Error("expected error from disabled cache, got nil")
	}
}

// ─── configHash ──────────────────────────────────────────────────────────────

func TestConfigHash_Deterministic(t *testing.T) {
	cfg := config.DiscoveryConfig{
		Enabled:           true,
		ScanShellConfigs:  true,
		InstallationPaths: []string{"/usr/bin", "/usr/local/bin"},
	}
	h1 := configHash(cfg)
	h2 := configHash(cfg)
	if h1 != h2 {
		t.Errorf("configHash not deterministic: %q != %q", h1, h2)
	}
}

func TestConfigHash_DifferentConfigs(t *testing.T) {
	cfg1 := config.DiscoveryConfig{
		Enabled:           true,
		InstallationPaths: []string{"/usr/bin"},
	}
	cfg2 := config.DiscoveryConfig{
		Enabled:           true,
		InstallationPaths: []string{"/usr/local/bin"},
	}
	if configHash(cfg1) == configHash(cfg2) {
		t.Error("expected different hashes for different configs, got same hash")
	}
}
