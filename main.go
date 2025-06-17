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

type Config struct {
	UseFlags              bool
	UseHotkeys            bool
	NoDiscovery           bool
	legacyEnableDiscovery bool
	EnableDiscovery       bool
	RefreshCache          bool
	ConfigPath            string
	ShowHelp              bool
}

type App struct {
	matcher *search.Matcher
	config  *Config
}

func parseFlags() *Config {
	cfg := &Config{}

	flag.BoolVar(&cfg.UseFlags, "flags", false, "Search in flags instead of commands")
	flag.BoolVar(&cfg.UseHotkeys, "hotkeys", false, "Search in hotkeys instead of commands")
	flag.BoolVar(&cfg.NoDiscovery, "no-discovery", false,
		"Disable system discovery (use only built-in commands)")
	flag.BoolVar(&cfg.RefreshCache, "refresh", false,
		"Force refresh of system discovery cache")
	flag.StringVar(&cfg.ConfigPath, "config", "",
		"Path to JSON configuration file")
	flag.BoolVar(&cfg.ShowHelp, "help", false, "Show this help message")
	flag.BoolVar(&cfg.legacyEnableDiscovery, "enable-discovery", false,
		"DEPRECATED: Discovery is now enabled by default when paths are configured")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "what-cmd discovers and searches terminal commands from built-in knowledge and configured paths.\n\n")
		fmt.Fprintf(os.Stderr, "OPTIONS:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nMODES:\n")
		fmt.Fprintf(os.Stderr, "  Default:        Uses built-in commands + configured installation paths\n")
		fmt.Fprintf(os.Stderr, "  --no-discovery: Uses only built-in commands (safe mode)\n")
		fmt.Fprintf(os.Stderr, "\nEXAMPLES:\n")
		fmt.Fprintf(os.Stderr, "  %s                              # Discover from configured paths\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s --no-discovery               # Built-in commands only\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s --flags                      # Search flags instead\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s --config=./config.json       # Custom configuration\n", os.Args[0])
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
	return nil
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

func performSystemDiscovery(items *[]models.Item, appConfig *config.Config, enableDiscovery, forceRefresh bool) error {
	if !enableDiscovery {
		return nil
	}

	effectivePaths := appConfig.Discovery.GetEffectivePaths()
	if len(effectivePaths) == 0 {
		fmt.Fprintf(os.Stderr, "No installation paths configured - skipping discovery\n")
		return nil
	}

	scanner, err := discovery.NewSystemScanner(appConfig)
	if err != nil {
		return fmt.Errorf("failed to create scanner: %w", err)
	}

	if forceRefresh {
		if err := scanner.ClearCache(); err != nil {
			fmt.Fprintf(os.Stderr, "Warning: Failed to clear cache: %v\n", err)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	discoveredItems, err := scanner.DiscoverItems(ctx)
	if err != nil {
		return fmt.Errorf("discovery failed: %w", err)
	}

	*items = append(*items, discoveredItems...)
	fmt.Fprintf(os.Stderr, "Discovery completed: found %d system items\n", len(discoveredItems))

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

	appConfig, err := config.LoadConfig(cfg.ConfigPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load configuration: %v\n", err)
		os.Exit(1)
	}

	enableDiscovery := !cfg.NoDiscovery && len(appConfig.Discovery.GetEffectivePaths()) > 0

	if enableDiscovery {
		fmt.Fprintf(os.Stderr, "DISCOVERY MODE: Scanning configured installation paths\n")
	} else if cfg.NoDiscovery {
		fmt.Fprintf(os.Stderr, "SAFE MODE: Discovery disabled - using built-in commands only\n")
	} else {
		fmt.Fprintf(os.Stderr, "BUILT-IN MODE: No installation paths configured\n")
	}

	items := loadBuiltinItems(cfg)
	fmt.Fprintf(os.Stderr, "Loaded %d built-in items\n", len(items))

	if enableDiscovery {
		if err := performSystemDiscovery(&items, appConfig, true, cfg.RefreshCache); err != nil {
			fmt.Fprintf(os.Stderr, "Discovery warning: %v\n", err)
			fmt.Fprintf(os.Stderr, "Continuing with built-in commands only\n")
		}
	}

	matcher := search.NewMatcher(items)

	app := NewApp(matcher, cfg)
	if err := app.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Application error: %v\n", err)
		os.Exit(1)
	}
}
