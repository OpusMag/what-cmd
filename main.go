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
}

type App struct {
	matcher *search.Matcher
	config  *Config
}

func main() {
	cfg := parseFlags()

	appConfig, err := config.LoadConfig(cfg.ConfigPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load configuration: %v\n", err)
		os.Exit(1)
	}

	operationMode := determineOperationMode(cfg)
	logOperationMode(operationMode)

	configureDiscovery(appConfig, operationMode)

	items := loadBuiltinItems(cfg)

	if appConfig.Discovery.Enabled {
		if err := performSystemDiscovery(&items, appConfig, operationMode, cfg.RefreshCache); err != nil {
			fmt.Fprintf(os.Stderr, "System discovery failed safely: %v\n", err)
		}
	}

	app := NewApp(search.NewMatcher(items), cfg)
	if err := app.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Application error: %v\n", err)
		os.Exit(1)
	}
}

func parseFlags() *Config {
	cfg := &Config{}

	flag.BoolVar(&cfg.UseFlags, "flags", false, "search in flags instead of commands")
	flag.BoolVar(&cfg.UseHotkeys, "hotkeys", false, "search in hotkeys instead of commands")

	flag.BoolVar(&cfg.EnableDiscovery, "enable-discovery", false, "enable system discovery with safety controls")
	flag.BoolVar(&cfg.UnsafeMode, "unsafe", false, "enable system discovery with minimal safety (NOT RECOMMENDED)")

	flag.BoolVar(&cfg.SafeMode, "safe", true, "run in safe mode (disables system discovery)")
	flag.BoolVar(&cfg.NoDiscovery, "no-discovery", false, "explicitly disable system discovery")
	flag.BoolVar(&cfg.EmergencyMode, "emergency", false, "emergency mode - complete operational shutdown")

	flag.BoolVar(&cfg.RefreshCache, "refresh", false, "force refresh of system discovery cache")
	flag.StringVar(&cfg.ConfigPath, "config", "", "path to configuration file")

	flag.Parse()

	if err := validateFlags(cfg); err != nil {
		fmt.Fprintf(os.Stderr, "Invalid flag combination: %v\n", err)
		flag.Usage()
		os.Exit(1)
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
		return fmt.Errorf("cannot use -flags and -hotkeys together")
	}

	discoveryFlags := 0
	if cfg.EnableDiscovery {
		discoveryFlags++
	}
	if cfg.UnsafeMode {
		discoveryFlags++
	}
	if discoveryFlags > 1 {
		return fmt.Errorf("cannot use -enable-discovery and -unsafe together")
	}

	if cfg.EmergencyMode && (cfg.EnableDiscovery || cfg.UnsafeMode) {
		fmt.Fprintf(os.Stderr, "WARNING: Emergency mode overrides discovery flags\n")
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
		fmt.Fprintf(os.Stderr, "SAFE MODE: Using built-in commands only\n")
	case DiscoveryMode:
		fmt.Fprintf(os.Stderr, "DISCOVERY MODE: System discovery enabled with comprehensive safety controls\n")
	case UnsafeMode:
		fmt.Fprintf(os.Stderr, "WARNING: Unsafe mode enabled - minimal safety controls active\n")
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
	scanner, err := createScanner(cfg, mode)
	if err != nil {
		return fmt.Errorf("failed to create scanner: %w", err)
	}

	if forceRefresh {
		if err := clearCache(scanner); err != nil {
			fmt.Fprintf(os.Stderr, "Warning: Failed to clear cache: %v\n", err)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	discoveredItems, err := scanner.DiscoverItems(ctx)
	if err != nil {
		return fmt.Errorf("discovery execution failed: %w", err)
	}

	*items = append(*items, discoveredItems...)
	fmt.Fprintf(os.Stderr, "Discovery completed successfully: found %d system items\n", len(discoveredItems))

	return nil
}

func createScanner(cfg *config.Config, mode OperationMode) (interface {
	DiscoverItems(context.Context) ([]models.Item, error)
}, error) {
	switch mode {
	case DiscoveryMode:
		return discovery.NewEmergencyScanner(cfg)
	case UnsafeMode:
		return discovery.NewSystemScanner(cfg)
	default:
		return nil, fmt.Errorf("invalid operation mode for system discovery: %s", mode)
	}
}

func clearCache(scanner interface{}) error {
	type CacheClearer interface {
		ClearCache() error
	}

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
