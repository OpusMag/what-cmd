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

	WindowsSafeMode     bool     `json:"windows_safe_mode"`
	WindowsAllowedPaths []string `json:"windows_allowed_paths"`
	WindowsBlockedPaths []string `json:"windows_blocked_paths"`
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
			ScanPATH:             getDefaultScanPATH(),
			ScanCommonPaths:      getDefaultScanCommonPaths(),
			ScanShellConfigs:     true,
			CustomPaths:          getDefaultCustomPaths(),
			CustomConfigFiles:    getDefaultConfigFiles(),
			ExcludePatterns:      getDefaultExcludePatterns(),
			MaxExecutables:       getDefaultMaxExecutables(),
			RefreshIntervalHours: 24,
			WindowsSafeMode:      getDefaultWindowsSafeMode(),
			WindowsAllowedPaths:  getDefaultWindowsAllowedPaths(),
			WindowsBlockedPaths:  getDefaultWindowsBlockedPaths(),
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

func getDefaultScanPATH() bool {
	return runtime.GOOS != "windows"
}

func getDefaultScanCommonPaths() bool {
	return runtime.GOOS != "windows"
}

func getDefaultMaxExecutables() int {
	if runtime.GOOS == "windows" {
		return 100
	}
	return 1000
}

func getDefaultWindowsSafeMode() bool {
	return runtime.GOOS == "windows"
}

func getDefaultCustomPaths() []string {
	switch runtime.GOOS {
	case "windows":
		return []string{
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

func getDefaultWindowsAllowedPaths() []string {
	if runtime.GOOS != "windows" {
		return []string{}
	}
	return []string{
		`C:\ProgramData\chocolatey\bin`,
		// more can be added via config.json:
	}
}

func getDefaultWindowsBlockedPaths() []string {
	if runtime.GOOS != "windows" {
		return []string{}
	}
	return []string{
		`C:\Windows`,
		`C:\Windows\System32`,
		`C:\Windows\SysWOW64`,
		`C:\Windows\WinSxS`,
		`C:\Program Files\Windows`,
		`C:\Program Files (x86)\Windows`,
		`C:\ProgramData\Microsoft`,
		`C:\Windows\Microsoft.NET`,
		`C:\Windows\assembly`,
		`C:\Windows\servicing`,
		`C:\Windows\SoftwareUpdate`,
		`C:\Windows\security`,
		`C:\Windows\SystemApps`,
		`C:\Program Files\WindowsApps`,
		`C:\Windows\system32\wbem`,
		`C:\Windows\system32\WindowsPowerShell`,
		`C:\Windows\system32\OpenSSH`,
		`C:\Windows\system32\config`,
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
	default:
		return []string{
			filepath.Join(homeDir, ".bashrc"),
			filepath.Join(homeDir, ".bash_profile"),
			filepath.Join(homeDir, ".zshrc"),
			filepath.Join(homeDir, ".profile"),
			filepath.Join(homeDir, ".bash_aliases"),
			filepath.Join(homeDir, ".gitconfig"),
			filepath.Join(homeDir, ".config", "fish", "config.fish"),
			"/etc/bash.bashrc",
			"/etc/bashrc",
			"/etc/profile",
			"/etc/zsh/zshrc",
		}
	}
}

func getDefaultExcludePatterns() []string {
	patterns := []string{
		"*.txt", "*.md", "*.log", "*.json", "*.xml", "*.yml", "*.yaml",
		"test*", "*test", "*_test", ".*",
		"README*", "LICENSE*", "CHANGELOG*",
	}

	if runtime.GOOS == "windows" {
		patterns = append(patterns, "*.dll", "*.pdb", "*.lib", "*.obj")
	} else {
		patterns = append(patterns, "*.so", "*.so.*", "*.dylib", "*.a", "*.o")
	}

	return patterns
}

func LoadConfig(configPath string) (*Config, error) {
	if configPath == "" {
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

	if runtime.GOOS == "windows" {
		if config.Discovery.WindowsBlockedPaths == nil {
			config.Discovery.WindowsBlockedPaths = getDefaultWindowsBlockedPaths()
		}
		if config.Discovery.WindowsAllowedPaths == nil {
			config.Discovery.WindowsAllowedPaths = getDefaultWindowsAllowedPaths()
		}
	}

	return &config, nil
}
