package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
)

type Config struct {
	Discovery DiscoveryConfig `json:"discovery"`
	UI        UIConfig        `json:"ui"`
	Cache     CacheConfig     `json:"cache"`
}

type DiscoveryConfig struct {
	Enabled              bool     `json:"enabled"`
	ScanPATH             bool     `json:"scan_path"`
	ScanCommonPaths      bool     `json:"scan_common_paths"`
	ScanShellConfigs     bool     `json:"scan_shell_configs"`
	CustomPaths          []string `json:"custom_paths"`
	CustomConfigFiles    []string `json:"custom_config_files"`
	ExcludePatterns      []string `json:"exclude_patterns"`
	MaxExecutables       int      `json:"max_executables"`
	RefreshIntervalHours int      `json:"refresh_interval_hours"`
}

type UIConfig struct {
	ShowSystemCommands  bool `json:"show_system_commands"`
	ShowBuiltinCommands bool `json:"show_builtin_commands"`
	ShowUserAliases     bool `json:"show_user_aliases"`
	GroupBySource       bool `json:"group_by_source"`
}

type CacheConfig struct {
	Enabled   bool   `json:"enabled"`
	Directory string `json:"directory"`
	MaxSizeMB int    `json:"max_size_mb"`
	TTLHours  int    `json:"ttl_hours"`
}

func DefaultConfig() *Config {
	homeDir, _ := os.UserHomeDir()

	config := &Config{
		Discovery: DiscoveryConfig{
			Enabled:              true,
			ScanPATH:             true,
			ScanCommonPaths:      true,
			ScanShellConfigs:     true,
			CustomPaths:          getDefaultCustomPaths(),
			CustomConfigFiles:    getDefaultConfigFiles(),
			ExcludePatterns:      getDefaultExcludePatterns(),
			MaxExecutables:       1000,
			RefreshIntervalHours: 24,
		},
		UI: UIConfig{
			ShowSystemCommands:  true,
			ShowBuiltinCommands: true,
			ShowUserAliases:     true,
			GroupBySource:       false,
		},
		Cache: CacheConfig{
			Enabled:   true,
			Directory: filepath.Join(homeDir, ".what-cmd", "cache"),
			MaxSizeMB: 50,
			TTLHours:  168,
		},
	}

	return config
}

func getDefaultCustomPaths() []string {
	switch runtime.GOOS {
	case "windows":
		return []string{
			`C:\Program Files\Git\usr\bin`,
			`C:\Program Files\PowerShell\7`,
			`C:\Windows\System32`,
			`C:\ProgramData\chocolatey\bin`,
		}
	case "darwin":
		return []string{
			"/usr/local/bin",
			"/opt/homebrew/bin",
			"/usr/local/sbin",
			"/opt/local/bin",
		}
	default:
		return []string{
			"/usr/local/bin",
			"/usr/bin",
			"/bin",
			"/snap/bin",
			"/var/lib/flatpak/exports/bin",
		}
	}
}

func getDefaultConfigFiles() []string {
	homeDir, _ := os.UserHomeDir()

	switch runtime.GOOS {
	case "windows":
		return []string{
			filepath.Join(homeDir, "Documents", "PowerShell", "profile.ps1"),
			filepath.Join(homeDir, ".gitconfig"),
		}
	case "darwin", "linux":
		return []string{
			filepath.Join(homeDir, ".bashrc"),
			filepath.Join(homeDir, ".bash_profile"),
			filepath.Join(homeDir, ".zshrc"),
			filepath.Join(homeDir, ".profile"),
			filepath.Join(homeDir, ".bash_aliases"),
			filepath.Join(homeDir, ".gitconfig"),
			filepath.Join(homeDir, ".config", "fish", "config.fish"),
		}
	default:
		return []string{}
	}
}

func getDefaultExcludePatterns() []string {
	return []string{
		"*.dll", "*.so", "*.dylib", // Libraries
		"*.txt", "*.md", "*.log", // Documentation
		"test*", "*test", // Test executables
		".*", // Hidden files
	}
}

func LoadConfig(configPath string) (*Config, error) {
	if configPath == "" {
		homeDir, _ := os.UserHomeDir()
		configPath = filepath.Join(homeDir, ".what-cmd", "config.json")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return DefaultConfig(), nil
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func (c *Config) SaveConfig(configPath string) error {
	if configPath == "" {
		homeDir, _ := os.UserHomeDir()
		configPath = filepath.Join(homeDir, ".what-cmd", "config.json")
	}

	if err := os.MkdirAll(filepath.Dir(configPath), 0755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(configPath, data, 0644)
}
