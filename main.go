package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"what-cmd/commands"
	"what-cmd/flags"
	"what-cmd/hotkeys"
	"what-cmd/internal/config"
	"what-cmd/internal/discovery"
	"what-cmd/internal/models"
	"what-cmd/internal/search"
	"what-cmd/internal/ui"
)

type OperationMode int

const (
	SafeMode OperationMode = iota
	DiscoveryMode
	UnsafeMode
	EmergencyMode
)

func (m OperationMode) String() string {
	switch m {
	case SafeMode:
		return "Safe"
	case DiscoveryMode:
		return "Discovery"
	case UnsafeMode:
		return "Unsafe"
	case EmergencyMode:
		return "Emergency"
	default:
		return "Unknown"
	}
}

type Config struct {
	UseFlags        bool
	UseHotkeys      bool
	EnableDiscovery bool
	UnsafeMode      bool
	SafeMode        bool
	NoDiscovery     bool
	EmergencyMode   bool
	RefreshCache    bool
	ConfigPath      string
	ShowHelp        bool
}

type App struct {
	matcher *search.Matcher
	config  *Config
}

type Scanner interface {
	DiscoverItems(context.Context) ([]models.Item, error)
}

type CacheClearer interface {
	ClearCache() error
}

func parseFlags() *Config {
	cfg := &Config{}

	flag.BoolVar(&cfg.UseFlags, "flags", false, "Search in flags instead of commands")
	flag.BoolVar(&cfg.UseHotkeys, "hotkeys", false, "Search in hotkeys instead of commands")

	flag.BoolVar(&cfg.EnableDiscovery, "enable-discovery", false,
		"Enable system discovery with comprehensive safety controls")
	flag.BoolVar(&cfg.UnsafeMode, "unsafe", false,
		"Enable system discovery with minimal safety controls (NOT RECOMMENDED)")

	flag.BoolVar(&cfg.SafeMode, "safe", false,
		"Explicitly run in safe mode - disables all system discovery")
	flag.BoolVar(&cfg.NoDiscovery, "no-discovery", false,
		"Explicitly disable system discovery")
	flag.BoolVar(&cfg.EmergencyMode, "emergency", false,
		"Emergency mode - complete operational shutdown")

	flag.BoolVar(&cfg.RefreshCache, "refresh", false,
		"Force refresh of system discovery cache")
	flag.StringVar(&cfg.ConfigPath, "config", "",
		"Path to JSON configuration file")

	flag.BoolVar(&cfg.ShowHelp, "help", false, "Show this help message")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "what-cmd is a command-line tool for discovering and searching terminal commands.\n\n")
		fmt.Fprintf(os.Stderr, "OPTIONS:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nSECURITY MODES:\n")
		fmt.Fprintf(os.Stderr, "  Default (safe):     Uses only built-in commands (RECOMMENDED)\n")
		fmt.Fprintf(os.Stderr, "  --enable-discovery: Safe system discovery with comprehensive controls\n")
		fmt.Fprintf(os.Stderr, "  --unsafe:           Minimal safety controls (NOT RECOMMENDED)\n")
		fmt.Fprintf(os.Stderr, "  --emergency:        Complete operational shutdown\n")
		fmt.Fprintf(os.Stderr, "\nEXAMPLES:\n")
		fmt.Fprintf(os.Stderr, "  %s                              # Safe mode (default)\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s --enable-discovery           # Safe system discovery\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s --flags                      # Search flags instead of commands\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s --config=./config.json       # Use custom configuration\n", os.Args[0])
	}

	flag.Parse()

	if cfg.ShowHelp {
		flag.Usage()
		os.Exit(0)
	}

	return cfg
}

func validateFlags(cfg *Config) error {
	contentFlags := 0
	if cfg.UseFlags {
		contentFlags++
	}
	if cfg.UseHotkeys {
		contentFlags++
	}
	if contentFlags > 1 {
		return fmt.Errorf("cannot use --flags and --hotkeys together")
	}

	discoveryFlags := 0
	if cfg.EnableDiscovery {
		discoveryFlags++
	}
	if cfg.UnsafeMode {
		discoveryFlags++
	}
	if discoveryFlags > 1 {
		return fmt.Errorf("cannot use --enable-discovery and --unsafe together")
	}

	if cfg.EmergencyMode && (cfg.EnableDiscovery || cfg.UnsafeMode) {
		fmt.Fprintf(os.Stderr, "WARNING: Emergency mode overrides all discovery flags\n")
	}

	return nil
}

func determineOperationMode(cfg *Config) OperationMode {
	if cfg.EmergencyMode {
		return EmergencyMode
	}

	if cfg.SafeMode || cfg.NoDiscovery {
		return SafeMode
	}

	if cfg.UnsafeMode {
		return UnsafeMode
	}

	if cfg.EnableDiscovery {
		return DiscoveryMode
	}

	return SafeMode
}

func logOperationMode(mode OperationMode) {
	switch mode {
	case SafeMode:
		fmt.Fprintf(os.Stderr, "SAFE MODE: Using built-in commands only (maximum security)\n")
	case DiscoveryMode:
		fmt.Fprintf(os.Stderr, "DISCOVERY MODE: System discovery enabled with comprehensive safety controls\n")
	case UnsafeMode:
		fmt.Fprintf(os.Stderr, "UNSAFE MODE: Minimal safety controls active - use with caution\n")
	case EmergencyMode:
		fmt.Fprintf(os.Stderr, "EMERGENCY MODE: All system interaction disabled\n")
	}
}

func configureDiscovery(appConfig *config.Config, mode OperationMode) {
	switch mode {
	case SafeMode, EmergencyMode:
		appConfig.Discovery.Enabled = false
	case DiscoveryMode, UnsafeMode:
		appConfig.Discovery.Enabled = true
		if mode == DiscoveryMode {
			appConfig.Discovery.ScanPATH = true
			appConfig.Discovery.ScanCommonPaths = true
			appConfig.Discovery.ScanShellConfigs = true
			appConfig.Discovery.MaxExecutables = 1000
			appConfig.Discovery.RefreshIntervalHours = 24
		}
	}
}

func loadBuiltinItems(cfg *Config) []models.Item {
	switch {
	case cfg.UseFlags:
		return search.ConvertFlags(flags.Words)
	case cfg.UseHotkeys:
		return search.ConvertHotkeys(hotkeys.Words)
	default:
		return search.ConvertCommands(commands.Words)
	}
}

func performSystemDiscovery(items *[]models.Item, cfg *config.Config, mode OperationMode, forceRefresh bool) error {
	if mode == SafeMode || mode == EmergencyMode {
		return nil
	}

	scanner, err := createScanner(cfg, mode)
	if err != nil {
		return fmt.Errorf("failed to create scanner: %w", err)
	}

	if forceRefresh {
		if err := clearCache(scanner); err != nil {
			fmt.Fprintf(os.Stderr, "Warning: Failed to clear cache: %v\n", err)
		}
	}

	timeout := 30 * time.Second
	if mode == UnsafeMode {
		timeout = 60 * time.Second
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	discoveredItems, err := scanner.DiscoverItems(ctx)
	if err != nil {
		return fmt.Errorf("discovery execution failed: %w", err)
	}

	*items = append(*items, discoveredItems...)
	fmt.Fprintf(os.Stderr, "Discovery completed successfully: found %d system items\n", len(discoveredItems))

	return nil
}

func createScanner(cfg *config.Config, mode OperationMode) (Scanner, error) {
	switch mode {
	case DiscoveryMode:
		return discovery.NewSystemScanner(cfg)
	case UnsafeMode:
		return discovery.NewSystemScanner(cfg)
	case EmergencyMode:
		return discovery.NewEmergencyScanner(cfg)
	case SafeMode:
		return nil, fmt.Errorf("safe mode does not support system discovery")
	default:
		return nil, fmt.Errorf("invalid operation mode for system discovery: %s", mode)
	}
}

func clearCache(scanner Scanner) error {
	if clearer, ok := scanner.(CacheClearer); ok {
		return clearer.ClearCache()
	}
	return nil
}

func NewApp(matcher *search.Matcher, cfg *Config) *App {
	return &App{
		matcher: matcher,
		config:  cfg,
	}
}

func (a *App) Run() error {
	terminalUI, err := ui.NewTerminalUI(a.matcher)
	if err != nil {
		return fmt.Errorf("failed to initialize terminal UI: %w", err)
	}

	return terminalUI.Run()
}

func main() {
	cfg := parseFlags()

	if err := validateFlags(cfg); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		flag.Usage()
		os.Exit(1)
	}

	mode := determineOperationMode(cfg)
	logOperationMode(mode)

	appConfig, err := config.LoadConfig(cfg.ConfigPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load configuration: %v\n", err)
		os.Exit(1)
	}

	configureDiscovery(appConfig, mode)

	items := loadBuiltinItems(cfg)
	fmt.Fprintf(os.Stderr, "📚 Loaded %d built-in items\n", len(items))

	if mode == DiscoveryMode || mode == UnsafeMode {
		if err := performSystemDiscovery(&items, appConfig, mode, cfg.RefreshCache); err != nil {
			fmt.Fprintf(os.Stderr, "Discovery failed: %v\n", err)
		}
	}

	matcher := search.NewMatcher(items)

	app := NewApp(matcher, cfg)
	if err := app.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Application error: %v\n", err)
		os.Exit(1)
	}
}
