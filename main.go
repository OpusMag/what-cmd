package main

import (
	"flag"
	"fmt"
	"os"

	"what-cmd/commands"
	"what-cmd/flags"
	"what-cmd/hotkeys"
	"what-cmd/internal/models"
	"what-cmd/internal/search"
	"what-cmd/internal/ui"
)

func main() {
	useFlags := flag.Bool("flags", false, "search in flags instead of commands")
	useHotkeys := flag.Bool("hotkeys", false, "search in hotkeys instead of commands")
	flag.Parse()

	var items []models.Item
	switch {
	case *useFlags:
		items = search.ConvertFlags(flags.Words)
	case *useHotkeys:
		items = search.ConvertHotkeys(hotkeys.Words)
	default:
		items = search.ConvertCommands(commands.Words)
	}

	matcher := search.NewMatcher(items)

	app := NewApp(matcher)
	if err := app.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
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
