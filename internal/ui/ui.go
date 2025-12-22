package ui

import (
	"fmt"
	"strings"
	"what-cmd/flags"
	"what-cmd/internal/models"
	"what-cmd/internal/search"

	"github.com/gdamore/tcell/v2"
)

type TerminalUI struct {
	screen  tcell.Screen
	matcher *search.Matcher
	state   *UIState
}

type UIState struct {
	userInput      []rune
	selectedIndex  int
	scrollPosition int
	inputChanged   bool
}

func NewTerminalUI(matcher *search.Matcher) (*TerminalUI, error) {
	screen, err := tcell.NewScreen()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize tcell: %w", err)
	}

	if err := screen.Init(); err != nil {
		screen.Fini()
		return nil, fmt.Errorf("failed to initialize screen: %w", err)
	}

	return &TerminalUI{
		screen:  screen,
		matcher: matcher,
		state: &UIState{
			userInput:      make([]rune, 0),
			selectedIndex:  0,
			scrollPosition: 0,
			inputChanged:   false,
		},
	}, nil
}

func (ui *TerminalUI) Run() error {
	defer ui.screen.Fini()

	for {
		if err := ui.render(); err != nil {
			return fmt.Errorf("render error: %w", err)
		}

		event := ui.screen.PollEvent()
		action, err := ui.handleEvent(event)
		if err != nil {
			return fmt.Errorf("event handling error: %w", err)
		}

		switch action {
		case ActionExit:
			return nil
		case ActionSelect:
			return ui.handleSelection()
		}
	}
}

type EventAction int

const (
	ActionContinue EventAction = iota
	ActionExit
	ActionSelect
)

func (ui *TerminalUI) handleEvent(event tcell.Event) (EventAction, error) {
	switch ev := event.(type) {
	case *tcell.EventKey:
		return ui.handleKeyEvent(ev), nil
	case *tcell.EventResize:
		ui.screen.Sync()
		return ActionContinue, nil
	case *tcell.EventError:
		return ActionExit, fmt.Errorf("terminal error: %w", fmt.Errorf("%s", ev.Error()))
	default:
		return ActionContinue, nil
	}
}

func (ui *TerminalUI) handleKeyEvent(ev *tcell.EventKey) EventAction {
	switch ev.Key() {
	case tcell.KeyEscape, tcell.KeyCtrlC:
		return ActionExit
	case tcell.KeyEnter:
		if strings.EqualFold(strings.TrimSpace(string(ui.state.userInput)), "exit") {
			return ActionExit
		}
		return ActionSelect
	case tcell.KeyUp, tcell.KeyDown:
		ui.handleScrollInput(ev)
	case tcell.KeyBackspace, tcell.KeyBackspace2:
		ui.handleBackspace()
	case tcell.KeyCtrlU:
		ui.clearInput()
	default:
		if ev.Rune() != 0 {
			ui.handleTextInput(ev.Rune())
		}
	}
	return ActionContinue
}

func (ui *TerminalUI) handleScrollInput(event *tcell.EventKey) {
	filteredItems := ui.getFilteredItems()
	_, height := ui.screen.Size()
	cmdWindowHeight := height - 10

	switch event.Key() {
	case tcell.KeyUp:
		if ui.state.selectedIndex > 0 {
			ui.state.selectedIndex--
			if ui.state.selectedIndex < ui.state.scrollPosition {
				ui.state.scrollPosition--
			}
		}
	case tcell.KeyDown:
		if ui.state.selectedIndex < len(filteredItems)-1 {
			ui.state.selectedIndex++
			if ui.state.selectedIndex >= ui.state.scrollPosition+cmdWindowHeight-2 {
				ui.state.scrollPosition++
			}
		}
	}
}

func (ui *TerminalUI) handleBackspace() {
	if len(ui.state.userInput) > 0 {
		ui.state.userInput = ui.state.userInput[:len(ui.state.userInput)-1]
		ui.state.inputChanged = true
	}
}

func (ui *TerminalUI) handleTextInput(r rune) {
	ui.state.userInput = append(ui.state.userInput, r)
	ui.state.inputChanged = true
}

func (ui *TerminalUI) getFilteredItems() []models.Item {
	query := string(ui.state.userInput)
	return ui.matcher.Search(query)
}

func (ui *TerminalUI) handleSelection() error {
	filteredItems := ui.getFilteredItems()
	if len(filteredItems) == 0 || ui.state.selectedIndex >= len(filteredItems) {
		return nil
	}

	selected := filteredItems[ui.state.selectedIndex]
	ui.screen.Fini()

	fmt.Printf("%s: %s\n", selected.Name, selected.Description)

	if flagsList := ui.getFlagsForCommand(selected.Name); len(flagsList) > 0 {
		for _, flag := range flagsList {
			fmt.Printf("  %s: %s\n", flag.Name, flag.Description)
		}
	}

	return nil
}

func (ui *TerminalUI) render() error {
	ui.screen.Clear()

	width, height := ui.screen.Size()

	if ui.state.inputChanged {
		ui.state.selectedIndex = 0
		ui.state.scrollPosition = 0
		ui.state.inputChanged = false
	}

	cmdWindowHeight := height - 10
	cmdWindowWidth := width * 2 / 10
	descWindowWidth := width - cmdWindowWidth - 2

	styles := ui.getStyles()

	ui.drawBorder(0, 0, cmdWindowWidth, cmdWindowHeight, styles.border)
	ui.drawBorder(cmdWindowWidth+1, 0, width-1, cmdWindowHeight, styles.border)

	ui.drawText("Commands", 1, 0, styles.header)
	ui.drawText("Descriptions", cmdWindowWidth+2, 0, styles.header)

	ui.drawSearchArea(cmdWindowWidth, cmdWindowHeight, styles)

	ui.drawASCIIArt(width, cmdWindowWidth, descWindowWidth, cmdWindowHeight, styles)

	ui.drawFilteredItems(cmdWindowWidth, descWindowWidth, cmdWindowHeight, styles)

	ui.drawFlags(width, cmdWindowWidth, cmdWindowHeight, styles)

	ui.drawStatusBar(width, height, styles)

	ui.screen.Show()
	return nil
}

func (ui *TerminalUI) drawSearchArea(cmdWindowWidth, cmdWindowHeight int, styles UIStyles) {
	promptBoxHeight := 8
	promptBoxYStart := cmdWindowHeight + 1
	promptBoxXEnd := cmdWindowWidth
	promptBoxYEnd := promptBoxYStart + promptBoxHeight

	ui.drawBorder(0, promptBoxYStart, promptBoxXEnd, promptBoxYEnd, styles.border)
	ui.drawText("Search", 1, promptBoxYStart, styles.header)

	prompt := "Enter a command to search for (type 'exit' to quit): "
	wrappedPrompt := ui.wrapText(prompt, cmdWindowWidth-2)
	for i, line := range wrappedPrompt {
		for j, r := range line {
			ui.screen.SetContent(1+j, promptBoxYStart+1+i, r, nil, tcell.StyleDefault)
		}
	}

	userInputYStart := promptBoxYEnd - 1
	for i, r := range ui.state.userInput {
		ui.screen.SetContent(1+i, userInputYStart, r, nil, styles.prompt)
	}
}

func (ui *TerminalUI) drawASCIIArt(width, cmdWindowWidth, descWindowWidth, cmdWindowHeight int, styles UIStyles) {
	asciiArt := `
        __          ___    _       _______      _____ __  __ _____  
        \ \        / / |  | |   /\|__   __|    / ____|  \/  |  __ \ 
         \ \  /\  / /| |__| |  /  \  | |______| |    | \  / | |  | |
          \ \/  \/ / |  __  | / /\ \ | |______| |    | |\/| | |  | |
           \  /\  /  | |  | |/ ____ \| |      | |____| |  | | |__| |
            \/  \/   |_|  |_/_/    \_\_|       \_____|_|  |_|_____/ `

	asciiArtLines := strings.Split(asciiArt, "\n")
	asciiArtHeight := len(asciiArtLines)
	asciiArtWidth := 0
	for _, line := range asciiArtLines {
		if len(line) > asciiArtWidth {
			asciiArtWidth = len(line)
		}
	}

	asciiHeight := 8
	asciiBoxYEnd := cmdWindowHeight
	asciiBoxYStart := asciiBoxYEnd - asciiHeight
	asciiArtX := cmdWindowWidth + descWindowWidth - asciiArtWidth + 1
	asciiArtY := asciiBoxYStart - asciiArtHeight + 8

	for y, line := range asciiArtLines {
		for x, r := range line {
			ui.screen.SetContent(asciiArtX+x, asciiArtY+y, r, nil, styles.border)
		}
	}
}

func (ui *TerminalUI) drawFilteredItems(cmdWindowWidth, descWindowWidth, cmdWindowHeight int, styles UIStyles) {
	filteredItems := ui.getFilteredItems()

	for i := ui.state.scrollPosition; i < len(filteredItems) && i < ui.state.scrollPosition+cmdWindowHeight-1; i++ {
		item := filteredItems[i]
		style := tcell.StyleDefault
		if i == ui.state.selectedIndex {
			style = styles.highlight
		}

		wrappedName := ui.wrapText(item.Name, cmdWindowWidth-2)
		wrappedDescription := ui.wrapText(item.Description, descWindowWidth-2)

		currentLine := i - ui.state.scrollPosition + 1
		for _, line := range wrappedName {
			if currentLine < cmdWindowHeight-1 {
				for j, r := range line {
					ui.screen.SetContent(j+1, currentLine, r, nil, style)
				}
				currentLine++
			}
		}

		currentLine = i - ui.state.scrollPosition + 1
		for _, line := range wrappedDescription {
			if currentLine < cmdWindowHeight-1 {
				for j, r := range line {
					ui.screen.SetContent(cmdWindowWidth+2+j, currentLine, r, nil, style)
				}
				currentLine++
			}
		}
	}
}

func (ui *TerminalUI) drawFlags(width, cmdWindowWidth, cmdWindowHeight int, styles UIStyles) {
	filteredItems := ui.getFilteredItems()

	// Only draw flags box if filtered results are reasonable
	if len(filteredItems) < 10000 {
		flagsBoxHeight := 8
		flagsBoxXStart := cmdWindowWidth + 1
		flagsBoxXEnd := width - 1
		flagsBoxYStart := cmdWindowHeight + 1
		flagsBoxYEnd := flagsBoxYStart + flagsBoxHeight
		flagsBoxWidth := width - cmdWindowWidth - 2

		ui.drawBorder(flagsBoxXStart, flagsBoxYStart, flagsBoxXEnd, flagsBoxYEnd, styles.border)
		ui.drawText("Flags", flagsBoxXStart+1, flagsBoxYStart, styles.header)

		if ui.state.selectedIndex < len(filteredItems) {
			selectedCommand := filteredItems[ui.state.selectedIndex]
			flagsList := ui.getFlagsForCommand(selectedCommand.Name)
			for i, flag := range flagsList {
				if i < flagsBoxYEnd-flagsBoxYStart-1 {
					y := flagsBoxYStart + 1 + i

					if len(flag.Name) > 0 {
						ui.screen.SetContent(flagsBoxXStart+1, y, rune(flag.Name[0]), nil, styles.highlight)
						for j, r := range flag.Name[1:] {
							ui.screen.SetContent(flagsBoxXStart+2+j, y, r, nil, styles.highlight)
						}
					}

					wrappedDescription := ui.wrapText(flag.Description, flagsBoxWidth-2-len(flag.Name)-1)
					for k, line := range wrappedDescription {
						if y+k < flagsBoxYEnd {
							for j, r := range line {
								ui.screen.SetContent(flagsBoxXStart+2+len(flag.Name)+1+j, y+k, r, nil, styles.highlight)
							}
						}
					}
				}
			}
		}
	}
}

type UIStyles struct {
	normal    tcell.Style
	highlight tcell.Style
	border    tcell.Style
	header    tcell.Style
	prompt    tcell.Style
}

func (ui *TerminalUI) getStyles() UIStyles {
	return UIStyles{
		normal:    tcell.StyleDefault.Foreground(tcell.ColorWhite),
		highlight: tcell.StyleDefault.Foreground(tcell.ColorLightSkyBlue).Bold(true),
		border:    tcell.StyleDefault.Foreground(tcell.ColorTeal),
		header:    tcell.StyleDefault.Foreground(tcell.ColorWhite),
		prompt:    tcell.StyleDefault.Foreground(tcell.ColorRed).Bold(true),
	}
}

func (ui *TerminalUI) drawBorder(x1, y1, x2, y2 int, style tcell.Style) {
	for x := x1; x <= x2; x++ {
		ui.screen.SetContent(x, y1, tcell.RuneHLine, nil, style)
		ui.screen.SetContent(x, y2, tcell.RuneHLine, nil, style)
	}
	for y := y1; y <= y2; y++ {
		ui.screen.SetContent(x1, y, tcell.RuneVLine, nil, style)
		ui.screen.SetContent(x2, y, tcell.RuneVLine, nil, style)
	}
	ui.screen.SetContent(x1, y1, tcell.RuneULCorner, nil, style)
	ui.screen.SetContent(x2, y1, tcell.RuneURCorner, nil, style)
	ui.screen.SetContent(x1, y2, tcell.RuneLLCorner, nil, style)
	ui.screen.SetContent(x2, y2, tcell.RuneLRCorner, nil, style)
}

func (ui *TerminalUI) drawText(text string, x, y int, style tcell.Style) {
	for i, r := range text {
		ui.screen.SetContent(x+i, y, r, nil, style)
	}
}

func (ui *TerminalUI) getFlagsForCommand(commandName string) []flags.Flag {
	if commandFlag, exists := flags.Words[commandName]; exists {
		return []flags.Flag{commandFlag}
	}
	return nil
}

func (ui *TerminalUI) wrapText(text string, maxWidth int) []string {
	if maxWidth <= 0 {
		return []string{""}
	}
	var lines []string
	words := strings.Fields(text)
	if len(words) == 0 {
		return lines
	}

	currentLine := ""
	flush := func() {
		if currentLine != "" {
			lines = append(lines, currentLine)
			currentLine = ""
		}
	}

	for _, w := range words {
		if len(w) > maxWidth {
			flush()
			for start := 0; start < len(w); start += maxWidth {
				end := start + maxWidth
				if end > len(w) {
					end = len(w)
				}
				lines = append(lines, w[start:end])
			}
			continue
		}
		if currentLine == "" {
			currentLine = w
		} else if len(currentLine)+1+len(w) > maxWidth {
			flush()
			currentLine = w
		} else {
			currentLine += " " + w
		}
	}
	flush()
	return lines
}

func (ui *TerminalUI) clearInput() {
	ui.state.userInput = ui.state.userInput[:0]
	ui.state.inputChanged = true
}

func (ui *TerminalUI) drawStatusBar(width, height int, styles UIStyles) {
	msg := "↑/↓ to navigate • Enter to select • Esc to quit • Ctrl+U to clear"
	y := height - 1
	for i, r := range msg {
		if i+1 >= width {
			break
		}
		ui.screen.SetContent(1+i, y, r, nil, styles.header)
	}
}
