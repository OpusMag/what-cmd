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

func main() {
	useFlags := flag.Bool("flags", false, "search in flags instead of commands")
	useHotkeys := flag.Bool("hotkeys", false, "search in hotkeys instead of commands")
	disableDiscovery := flag.Bool("no-discovery", false, "disable system discovery")
	safeMode := flag.Bool("safe", true, "run in safe mode (recommended)")
	emergencyMode := flag.Bool("emergency", false, "emergency mode - disable all discovery")
	refreshCache := flag.Bool("refresh", false, "force refresh of system discovery cache")
	configPath := flag.String("config", "", "path to configuration file")
	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load configuration: %v\n", err)
		os.Exit(1)
	}

	if *disableDiscovery || *safeMode || *emergencyMode {
		cfg.Discovery.Enabled = false
		fmt.Fprintf(os.Stderr, "SAFETY: System discovery disabled\n")
	}

	var items []models.Item

	switch {
	case *useFlags:
		items = search.ConvertFlags(flags.Words)
	case *useHotkeys:
		items = search.ConvertHotkeys(hotkeys.Words)
	default:
		items = search.ConvertCommands(commands.Words)
	}

	if cfg.Discovery.Enabled && !*emergencyMode {
		fmt.Fprintf(os.Stderr, "WARNING: System discovery enabled - use --safe to disable\n")

		if err := discoverSystemItemsSafely(&items, cfg, *refreshCache); err != nil {
			fmt.Fprintf(os.Stderr, "System discovery failed safely: %v\n", err)
		}
	} else {
		fmt.Fprintf(os.Stderr, "SAFE MODE: Using built-in commands only\n")
	}

	matcher := search.NewMatcher(items)
	app := NewApp(matcher)

	if err := app.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func discoverSystemItemsSafely(items *[]models.Item, cfg *config.Config, forceRefresh bool) error {

	scanner, err := discovery.NewEmergencyScanner(cfg)
	if err != nil {
		return fmt.Errorf("failed to create emergency scanner: %w", err)
	}

	if forceRefresh {
		fmt.Fprintf(os.Stderr, "SAFETY: Cache refresh in safe mode\n")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	discoveredItems, err := scanner.DiscoverItems(ctx)
	if err != nil {
		return fmt.Errorf("safe discovery failed: %w", err)
	}

	fmt.Fprintf(os.Stderr, "SAFE DISCOVERY: Found %d items\n", len(discoveredItems))
	*items = append(*items, discoveredItems...)

	return nil
}

type App struct {
	matcher *search.Matcher
}

func NewApp(matcher *search.Matcher) *App {
	return &App{matcher: matcher}
}

func (a *App) Run() error {
	terminalUI, err := ui.NewTerminalUI(a.matcher)
	if err != nil {
		return fmt.Errorf("failed to create terminal UI: %w", err)
	}

	return terminalUI.Run()
}
