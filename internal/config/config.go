package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type Config struct {
	Discovery DiscoveryConfig `json:"discovery"`
	UI        UIConfig        `json:"ui"`
	Cache     CacheConfig     `json:"cache"`
}

type DiscoveryConfig struct {
	Enabled          bool `json:"enabled"`
	ScanShellConfigs bool `json:"scan_shell_configs"`

	InstallationPaths []string `json:"installation_paths"`
	UserPaths         []string `json:"user_paths"`

	// Legacy support (deprecated but available for backward compatibility)
	ScanPATH          bool     `json:"scan_path"`
	ScanCommonPaths   bool     `json:"scan_common_paths"`
	CustomPaths       []string `json:"custom_paths"`
	CustomConfigFiles []string `json:"custom_config_files"`

	ExcludePatterns      []string `json:"exclude_patterns"`
	MaxExecutables       int      `json:"max_executables"`
	RefreshIntervalHours int      `json:"refresh_interval_hours"`

	// Windows-specific security (maintained for compatibility)
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
			Enabled:           true,
			ScanShellConfigs:  true,
			InstallationPaths: getPlatformInstallationPaths(),
			UserPaths:         getUserDefaultPaths(),

			ScanPATH:          false,
			ScanCommonPaths:   false,
			CustomPaths:       []string{},
			CustomConfigFiles: getDefaultConfigFiles(),

			ExcludePatterns:      getDefaultExcludePatterns(),
			MaxExecutables:       getDefaultMaxExecutables(),
			RefreshIntervalHours: 24,

			WindowsSafeMode:     getDefaultWindowsSafeMode(),
			WindowsAllowedPaths: []string{},
			WindowsBlockedPaths: getDefaultWindowsBlockedPaths(),
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

func getPlatformInstallationPaths() []string {
	switch runtime.GOOS {
	case "windows":
		return []string{
			`C:\ProgramData\chocolatey\bin`,
			`C:\Program Files\Git\usr\bin`,
			`C:\Program Files\PowerShell\7`,
			`C:\Program Files\Go\bin`,
			`C:\Program Files\Python\Python311\Scripts`,
			`C:\Program Files\Python\Python312\Scripts`,
		}
	case "darwin":
		return []string{
			"/usr/local/bin",
			"/opt/homebrew/bin",
			"/usr/local/sbin",
			"/opt/local/bin",
			"/opt/homebrew/sbin",
		}
	default:
		return []string{
			"/usr/local/bin",
			"/usr/bin",
			"/bin",
			"/snap/bin",
			"/var/lib/flatpak/exports/bin",
			"/opt/bin",
			"usr/local/sbin",
			"/usr/sbin",
			"/sbin",
			"etc/bin",
			"etc/sbin",
		}
	}
}

func getUserDefaultPaths() []string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return []string{}
	}

	switch runtime.GOOS {
	case "windows":
		return []string{
			filepath.Join(homeDir, "bin"),
			filepath.Join(homeDir, "scripts"),
			filepath.Join(homeDir, "tools"),
		}
	default:
		return []string{
			filepath.Join(homeDir, ".local", "bin"),
			filepath.Join(homeDir, "bin"),
			filepath.Join(homeDir, ".cargo", "bin"),
			filepath.Join(homeDir, "go", "bin"),
			filepath.Join(homeDir, ".npm-global", "bin"),
		}
	}
}

func getDefaultMaxExecutables() int {
	if runtime.GOOS == "windows" {
		return 500
	}
	return 1000
}

func getDefaultWindowsSafeMode() bool {
	return runtime.GOOS == "windows"
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
		"*debug*", "*Debug*", "*DEBUG*",
		"*.tmp", "*.temp", "*.bak", "*.backup",
	}

	if runtime.GOOS == "windows" {
		patterns = append(patterns, "*.dll", "*.pdb", "*.lib", "*.obj", "*.exe.config")
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

	applyConfigDefaults(&config)

	migrateConfigToNewFormat(&config)

	return &config, nil
}

func applyConfigDefaults(config *Config) {
	if len(config.Discovery.InstallationPaths) == 0 {
		config.Discovery.InstallationPaths = getPlatformInstallationPaths()
	}

	if len(config.Discovery.UserPaths) == 0 {
		config.Discovery.UserPaths = getUserDefaultPaths()
	}

	if runtime.GOOS == "windows" {
		if config.Discovery.WindowsBlockedPaths == nil {
			config.Discovery.WindowsBlockedPaths = getDefaultWindowsBlockedPaths()
		}

		if !config.Discovery.WindowsSafeMode {
			config.Discovery.WindowsSafeMode = true
		}
	}
}

func migrateConfigToNewFormat(config *Config) {
	if len(config.Discovery.InstallationPaths) == 0 && len(config.Discovery.CustomPaths) > 0 {
		config.Discovery.InstallationPaths = append(config.Discovery.InstallationPaths, config.Discovery.CustomPaths...)
	}

	if runtime.GOOS == "windows" && len(config.Discovery.WindowsAllowedPaths) > 0 {
		existing := make(map[string]bool)
		for _, path := range config.Discovery.InstallationPaths {
			existing[strings.ToLower(path)] = true
		}

		for _, path := range config.Discovery.WindowsAllowedPaths {
			if !existing[strings.ToLower(path)] {
				config.Discovery.InstallationPaths = append(config.Discovery.InstallationPaths, path)
				existing[strings.ToLower(path)] = true
			}
		}
	}
}

func ValidateConfig(config *Config) error {
	if config.Discovery.MaxExecutables <= 0 {
		return fmt.Errorf("max_executables must be positive, got %d", config.Discovery.MaxExecutables)
	}

	if config.Discovery.RefreshIntervalHours <= 0 {
		return fmt.Errorf("refresh_interval_hours must be positive, got %d", config.Discovery.RefreshIntervalHours)
	}

	if config.Cache.MaxSizeMB <= 0 {
		return fmt.Errorf("cache max_size_mb must be positive, got %d", config.Cache.MaxSizeMB)
	}

	if config.Cache.TTLHours <= 0 {
		return fmt.Errorf("cache ttl_hours must be positive, got %d", config.Cache.TTLHours)
	}

	return nil
}

func (d *DiscoveryConfig) GetEffectivePaths() []string {
	var paths []string

	paths = append(paths, d.InstallationPaths...)

	paths = append(paths, d.UserPaths...)

	paths = append(paths, d.CustomPaths...)

	return removeDuplicatePaths(paths)
}

func removeDuplicatePaths(paths []string) []string {
	seen := make(map[string]bool)
	var unique []string

	for _, path := range paths {
		normalized := strings.ToLower(filepath.Clean(path))
		if !seen[normalized] {
			seen[normalized] = true
			unique = append(unique, path)
		}
	}

	return unique
}
