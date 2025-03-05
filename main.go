package main

import (
    "flag"
    "fmt"
    "os"
    "strings"
    "github.com/gdamore/tcell/v2"
    "what-cmd/commands"
    "what-cmd/flags"
    "what-cmd/hotkeys"
)

type KeyValuePair struct {
    Name        string
    Description string
}

// converting the flag map into a slice to allow uniform processing and ranking.
func convertMapToKeyValuePairs(m map[string]flags.Flag) []KeyValuePair {
    var keyValuePairs []KeyValuePair
    for k, v := range m {
        keyValuePairs = append(keyValuePairs, KeyValuePair{Name: k, Description: v.Description})
    }
    return keyValuePairs
}

// transforming command definitions to a uniform structure so our matching logic handles all items similarly.
func convertCommandsToKeyValuePairs(cmds []commands.Command) []KeyValuePair {
    var keyValuePairs []KeyValuePair
    for _, c := range cmds {
        keyValuePairs = append(keyValuePairs, KeyValuePair{Name: c.Name, Description: c.Description})
    }
    return keyValuePairs
}

// mapping hotkey definitions to a key-value pair form to integrate them into our matching algorithm.
func convertHotkeysToKeyValuePairs(cmds []hotkeys.Hotkey) []KeyValuePair {
    var keyValuePairs []KeyValuePair
    for _, c := range cmds {
        keyValuePairs = append(keyValuePairs, KeyValuePair{Name: c.Name, Description: c.Description})
    }
    return keyValuePairs
}

// using a weighted scoring approach to determine the most relevant command based on user input,
// ensuring that exact matches and near similarities are prioritized over less relevant items.
func findClosestMatch(input string, words []KeyValuePair) KeyValuePair {
    var bestMatch KeyValuePair
    highestScore := -1

    for _, word := range words {
        score := 0

        // Prioritize an exact match in the name to align precisely with what the user intends.
        if strings.EqualFold(word.Name, input) {
            score += 100
        } else {
            // Searching descriptions can capture additional context the name alone might miss.
            if strings.Contains(strings.ToLower(word.Description), strings.ToLower(input)) {
                score += 10
            }

            // Partial matches in the name still count but receive a lower weight.
            if strings.Contains(strings.ToLower(word.Name), strings.ToLower(input)) {
                score += 5
            }

            // Using the edit distance helps us favor commands with minimal differences.
            nameDistance := distanceAtoB(strings.ToLower(word.Name), strings.ToLower(input))
            descDistance := distanceAtoB(strings.ToLower(word.Description), strings.ToLower(input))
            score += max(0, 10-nameDistance)
            score += max(0, 5-descDistance) 
        }

        if score > highestScore {
            highestScore = score
            bestMatch = word
        }
    }

    return bestMatch
}

// implementing Levenshtein distance to measure similarity,
// because quantifying edit differences aids in recognizing near misses.
func distanceAtoB(str1, str2 string) int {
    len1, len2 := len(str1), len(str2)
    distanceMatrix := make([][]int, len1+1)
    for row := range distanceMatrix {
        distanceMatrix[row] = make([]int, len2+1)
    }
    // filling the matrix to capture incremental edit costs, ensuring a fair comparison between strings.
    for row := 0; row <= len1; row++ {
        for col := 0; col <= len2; col++ {
            if row == 0 {
                distanceMatrix[row][col] = col
            } else if col == 0 {
                distanceMatrix[row][col] = row
            } else if str1[row-1] == str2[col-1] {
                distanceMatrix[row][col] = distanceMatrix[row-1][col-1]
            } else {
                distanceMatrix[row][col] = 1 + findMin(
                    distanceMatrix[row-1][col],
                    distanceMatrix[row][col-1],
                    distanceMatrix[row-1][col-1],
                )
            }
        }
    }

    return distanceMatrix[len1][len2]
}

// choosing the minimum among three possible edit paths to favor the least disruptive change.
func findMin(a, b, c int) int {
    if a < b && a < c {
        return a
    } else if b < c {
        return b
    }
    return c
}

// selecting the larger of two values to ensure our scoring in the matching algorithm captures the dominant factor.
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

// drawing borders to separate interface sections because clear boundaries help users navigate the UI.
func drawBorder(screen tcell.Screen, x1, y1, x2, y2 int, style tcell.Style) {
    for x := x1; x <= x2; x++ {
        screen.SetContent(x, y1, tcell.RuneHLine, nil, style)
        screen.SetContent(x, y2, tcell.RuneHLine, nil, style)
    }
    for y := y1; y <= y2; y++ {
        screen.SetContent(x1, y, tcell.RuneVLine, nil, style)
        screen.SetContent(x2, y, tcell.RuneVLine, nil, style)
    }
    screen.SetContent(x1, y1, tcell.RuneULCorner, nil, style)
    screen.SetContent(x2, y1, tcell.RuneURCorner, nil, style)
    screen.SetContent(x1, y2, tcell.RuneLLCorner, nil, style)
    screen.SetContent(x2, y2, tcell.RuneLRCorner, nil, style)
}

func getFlagsForCommand(commandName string) []flags.Flag {
    if commandFlag, exists := flags.Words[commandName]; exists {
        return []flags.Flag{commandFlag}
    }
    return nil
}

func wrapText(text string, maxWidth int) []string {
    var lines []string
    words := strings.Fields(text)
    if len(words) == 0 {
        return lines
    }

    currentLine := words[0]
    for _, word := range words[1:] {
        if len(currentLine)+len(word)+1 > maxWidth {
            lines = append(lines, currentLine)
            currentLine = word
        } else {
            currentLine += " " + word
        }
    }
    lines = append(lines, currentLine)
    return lines
}

func handleScrollInput(event *tcell.EventKey, scrollPosition *int, selectedIndex *int, filteredWords []KeyValuePair, cmdWindowHeight int) {
    switch event.Key() {
    case tcell.KeyUp:
        if *selectedIndex > 0 {
            *selectedIndex--
            if *selectedIndex < *scrollPosition {
                *scrollPosition--
            }
        }
    case tcell.KeyDown:
        if *selectedIndex < len(filteredWords)-1 {
            *selectedIndex++
            if *selectedIndex >= *scrollPosition+cmdWindowHeight-2 {
                *scrollPosition++
            }
        }
    }
}

func main() {
    flag.CommandLine.Usage = func() {
        fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
        flag.PrintDefaults()
    }

    useFlags := flag.Bool("flags", false, "search in flags instead of commands")
    useHotkeys := flag.Bool("hotkeys", false, "search in hotkeys instead of commands")
    flag.Parse()

    var words []KeyValuePair
    if *useFlags {
        words = convertMapToKeyValuePairs(flags.Words)
    } else if *useHotkeys {
        words = convertHotkeysToKeyValuePairs(hotkeys.Words)
    } else {
        words = convertCommandsToKeyValuePairs(commands.Words)
    }

    // initializing the screen using tcell to gain platform-independent control of terminal rendering.
    screen, err := tcell.NewScreen()
    if err != nil {
        fmt.Println("Failed to initialize tcell:", err)
        os.Exit(1)
    }
    if err := screen.Init(); err != nil {
        fmt.Println("Failed to initialize screen:", err)
        os.Exit(1)
    }
    defer screen.Fini()

    var userInput []rune

    selectedIndex := 0

    inputChanged := false

    scrollPosition := 0

    // The main loop continuously refreshes the UI to represent the current state,
    // supporting a dynamic interface that reacts in real time.
    for {
        screen.Clear()

        width, height := screen.Size()

        cmdWindowHeight := height - 10
        cmdWindowWidth := width * 2 / 10
        descWindowWidth := width - cmdWindowWidth - 2

        whiteStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite)
        tealStyle := tcell.StyleDefault.Foreground(tcell.ColorTeal)
        highlightStyle := tcell.StyleDefault.Foreground(tcell.ColorLightSkyBlue).Bold(true)
        promptStyle := tcell.StyleDefault.Foreground(tcell.ColorRed).Bold(true)

        drawBorder(screen, 0, 0, cmdWindowWidth, cmdWindowHeight, tealStyle)
        drawBorder(screen, cmdWindowWidth+1, 0, width-1, cmdWindowHeight, tealStyle)

        commandWord := "Commands"
        for i, r := range commandWord {
            screen.SetContent(1+i, 0, r, nil, whiteStyle)
        }

        descriptionWord := "Descriptions"
        for i, r := range descriptionWord {
            screen.SetContent(cmdWindowWidth+2+i, 0, r, nil, whiteStyle)
        }

        promptBoxHeight := 8 
        promptBoxWidth := cmdWindowWidth 
        promptBoxXStart := 0 
        promptBoxYStart := cmdWindowHeight + 1
        promptBoxXEnd := promptBoxXStart + promptBoxWidth
        promptBoxYEnd := promptBoxYStart + promptBoxHeight

        drawBorder(screen, promptBoxXStart, promptBoxYStart, promptBoxXEnd, promptBoxYEnd, tealStyle)

        searchWord := "Search"
        for i, r := range searchWord {
            screen.SetContent(promptBoxXStart+1+i, promptBoxYStart, r, nil, whiteStyle)
        }

        prompt := "Enter a command to search for (type 'exit' to quit): "
        wrappedPrompt := wrapText(prompt, promptBoxWidth-2)
        for i, line := range wrappedPrompt {
            for j, r := range line {
                screen.SetContent(promptBoxXStart+1+j, promptBoxYStart+1+i, r, nil, tcell.StyleDefault)
            }
        }

        userInputYStart := promptBoxYEnd - 1
        for i, r := range userInput {
            screen.SetContent(promptBoxXStart+1+i, userInputYStart, r, nil, promptStyle)
        }

        inputStr := string(userInput)
        closest := findClosestMatch(inputStr, words)

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
                screen.SetContent(asciiArtX+x, asciiArtY+y, r, nil, tealStyle)
            }
        }

        var filteredWords []KeyValuePair
        for _, word := range words {
            if strings.Contains(strings.ToLower(word.Name), strings.ToLower(inputStr)) || strings.Contains(strings.ToLower(word.Description), strings.ToLower(inputStr)) {
                filteredWords = append(filteredWords, word)
            }
        }

        closest = findClosestMatch(inputStr, filteredWords)

        if len(filteredWords) > 0 && closest.Name != "" {
            for i, word := range filteredWords {
                if word.Name == closest.Name && word.Description == closest.Description {
                    filteredWords[0], filteredWords[i] = filteredWords[i], filteredWords[0]
                    break
                }
            }
        }

        if inputChanged {
            selectedIndex = 0
            scrollPosition = 0 
            inputChanged = false
        }


        for i := scrollPosition; i < len(filteredWords) && i < scrollPosition+cmdWindowHeight-1; i++ {
            word := filteredWords[i]
            style := tcell.StyleDefault
            if i == selectedIndex {
                style = highlightStyle
            }

            wrappedName := wrapText(word.Name, cmdWindowWidth-2)
            wrappedDescription := wrapText(word.Description, descWindowWidth-2)

            currentLine := i - scrollPosition + 1
            for _, line := range wrappedName {
                if currentLine < cmdWindowHeight-1 {
                    for j, r := range line {
                        screen.SetContent(j+1, currentLine, r, nil, style)
                    }
                    currentLine++
                }
            }

            currentLine = i - scrollPosition + 1
            for _, line := range wrappedDescription {
                if currentLine < cmdWindowHeight-1 {
                    for j, r := range line {
                        screen.SetContent(cmdWindowWidth+2+j, currentLine, r, nil, style)
                    }
                    currentLine++
                }
            }
        }

        // If the number of filtered search results is less than 10000, draw a box and display flags
        if len(filteredWords) < 10000 {
            flagsBoxWidth := descWindowWidth 
            flagsBoxHeight := 8 
            flagsBoxXStart := cmdWindowWidth + 1
            flagsBoxXEnd := width - 1 
            flagsBoxYStart := cmdWindowHeight + 1 
            flagsBoxYEnd := flagsBoxYStart + flagsBoxHeight

            drawBorder(screen, flagsBoxXStart, flagsBoxYStart, flagsBoxXEnd, flagsBoxYEnd, tealStyle)

            word := "Flags"
            for i, r := range word {
                screen.SetContent(flagsBoxXStart+1+i, flagsBoxYStart, r, nil, whiteStyle)
            }

            if selectedIndex < len(filteredWords) {
                selectedCommand := filteredWords[selectedIndex]
                flags := getFlagsForCommand(selectedCommand.Name)
                for i, flag := range flags {
                    if i < flagsBoxYEnd-flagsBoxYStart-1 {
                        y := flagsBoxYStart + 1 + i
                        screen.SetContent(flagsBoxXStart+1, y, rune(flag.Name[0]), nil, highlightStyle)
                        for j, r := range flag.Name[1:] {
                            screen.SetContent(flagsBoxXStart+2+j, y, r, nil, highlightStyle)
                        }

                        wrappedDescription := wrapText(flag.Description, flagsBoxWidth-2-len(flag.Name)-1)
                        for k, line := range wrappedDescription {
                            if y+k < flagsBoxYEnd {
                                for j, r := range line {
                                    screen.SetContent(flagsBoxXStart+2+len(flag.Name)+1+j, y+k, r, nil, highlightStyle)
                                }
                            }
                        }
                    }
                }
            }
        }

        // Flushing the changes to the screen so they are displayed correctly
        screen.Show()

        event := screen.PollEvent()
        switch ev := event.(type) {
        case *tcell.EventKey:
            switch ev.Key() {
            case tcell.KeyEscape, tcell.KeyCtrlC:
                return
            case tcell.KeyUp, tcell.KeyDown:
                handleScrollInput(ev, &scrollPosition, &selectedIndex, filteredWords, cmdWindowHeight)
            case tcell.KeyEnter:
                selectedCommand := filteredWords[selectedIndex]
                screen.Fini()
                fmt.Printf("%s: %s\n", selectedCommand.Name, selectedCommand.Description)
                flags := getFlagsForCommand(selectedCommand.Name)
                for _, flag := range flags {
                    fmt.Printf("  %s: %s\n", flag.Name, flag.Description)
                }
                return
            case tcell.KeyBackspace, tcell.KeyBackspace2:
                if len(userInput) > 0 {
                    userInput = userInput[:len(userInput)-1]
                    inputChanged = true
                }
            default:
                if ev.Rune() != 0 {
                    userInput = append(userInput, ev.Rune())
                    inputChanged = true
                }
            }
        case *tcell.EventResize:
            screen.Sync()
        case *tcell.EventError:
            fmt.Println("Error:", ev.Error())
            return
        }
    }
}