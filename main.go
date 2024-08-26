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

// KeyValuePair structure
type KeyValuePair struct {
    Name        string
    Description string
}

// ConvertMapToKeyValuePairs converts a map to a slice of KeyValuePair
func convertMapToKeyValuePairs(m map[string]flags.Flag) []KeyValuePair {
    var keyValuePairs []KeyValuePair
    for k, v := range m {
        keyValuePairs = append(keyValuePairs, KeyValuePair{Name: k, Description: v.Description})
    }
    return keyValuePairs
}

// Convert commands.Command to KeyValuePair
func convertCommandsToKeyValuePairs(cmds []commands.Command) []KeyValuePair {
    var keyValuePairs []KeyValuePair
    for _, c := range cmds {
        keyValuePairs = append(keyValuePairs, KeyValuePair{Name: c.Name, Description: c.Description})
    }
    return keyValuePairs
}

// Convert hotkeys.Hotkey to KeyValuePair
func convertHotkeysToKeyValuePairs(cmds []hotkeys.Hotkey) []KeyValuePair {
    var keyValuePairs []KeyValuePair
    for _, c := range cmds {
        keyValuePairs = append(keyValuePairs, KeyValuePair{Name: c.Name, Description: c.Description})
    }
    return keyValuePairs
}

// Find the closest match to the input
func findClosestMatch(input string, words []KeyValuePair) KeyValuePair {
    var bestMatch KeyValuePair
    highestScore := -1

    for _, word := range words {
        score := 0

        // Check if the input exactly matches the Name
        if strings.EqualFold(word.Name, input) {
            score += 100 // Highest weight for exact matches in the Name
        } else {
            // Check if the input is in the Description
            if strings.Contains(strings.ToLower(word.Description), strings.ToLower(input)) {
                score += 10 // Higher weight for matches in the Description
            }

            // Check if the input is in the Name
            if strings.Contains(strings.ToLower(word.Name), strings.ToLower(input)) {
                score += 5 // Lower weight for matches in the Name
            }

            // Calculate Levenshtein distance for partial matches
            nameDistance := distanceAtoB(strings.ToLower(word.Name), strings.ToLower(input))
            descDistance := distanceAtoB(strings.ToLower(word.Description), strings.ToLower(input))
            score += max(0, 10-nameDistance) // Higher score for closer matches
            score += max(0, 5-descDistance)   // Higher score for closer matches
        }

        // Update the best match if the current score is higher
        if score > highestScore {
            highestScore = score
            bestMatch = word
        }
    }

    return bestMatch
}

// Calculate the distance between two strings (using Levenshtein)
func distanceAtoB(str1, str2 string) int {
    len1, len2 := len(str1), len(str2)
    distanceMatrix := make([][]int, len1+1)
    for row := range distanceMatrix {
        distanceMatrix[row] = make([]int, len2+1)
    }

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

// Finds the minimum of three integers
func findMin(a, b, c int) int {
    if a < b && a < c {
        return a
    } else if b < c {
        return b
    }
    return c
}

// Finds the maximum of two integers
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

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

// Get the flags
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

func main() {
    // Setting the CLI name
    flag.CommandLine.Usage = func() {
        fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
        flag.PrintDefaults()
    }

    // Defining flags
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

    // Initializing tcell
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

    // Creating buffer for user input
    var userInput []rune

    // Add a variable to track the selected index
    selectedIndex := 0

    // Main loop
    for {
        // Clearing the screen
        screen.Clear()

        // Get terminal dimensions
        width, height := screen.Size()

        // Adjusting window dimensions
        cmdWindowHeight := height - 5
        cmdWindowWidth := width * 2 / 10 // 20% of the terminal width (reduced by 50%)
        descWindowWidth := width - cmdWindowWidth - 2

        // Define teal color style
        tealStyle := tcell.StyleDefault.Foreground(tcell.ColorTeal)
        highlightStyle := tcell.StyleDefault.Foreground(tcell.ColorLightSkyBlue).Bold(true)

        // Draw borders for the command and description windows
        drawBorder(screen, 0, 0, cmdWindowWidth, cmdWindowHeight, tealStyle)
        drawBorder(screen, cmdWindowWidth+1, 0, width-1, cmdWindowHeight, tealStyle)

        // Displaying the prompt at the bottom left
        prompt := "Enter a command to search for (type 'exit' to quit): "
        for i, r := range prompt {
            screen.SetContent(i, height-3, r, nil, tcell.StyleDefault)
        }

        // Displaying the user input at the bottom left
        for i, r := range userInput {
            screen.SetContent(i, height-2, r, nil, tcell.StyleDefault)
        }

        // Finding and displaying the closest match
        inputStr := string(userInput)
        closest := findClosestMatch(inputStr, words)

        // Define the ASCII art
        asciiArt := `
        __          ___    _       _______      _____ __  __ _____  
        \ \        / / |  | |   /\|__   __|    / ____|  \/  |  __ \ 
         \ \  /\  / /| |__| |  /  \  | |______| |    | \  / | |  | |
          \ \/  \/ / |  __  | / /\ \ | |______| |    | |\/| | |  | |
           \  /\  /  | |  | |/ ____ \| |      | |____| |  | | |__| |
            \/  \/   |_|  |_/_/    \_\_|       \_____|_|  |_|_____/ 
                                                                        
        `

        // Calculate the starting position for the ASCII art
        asciiArtLines := strings.Split(asciiArt, "\n")
        asciiArtHeight := len(asciiArtLines)
        asciiArtWidth := 0
        for _, line := range asciiArtLines {
            if len(line) > asciiArtWidth {
                asciiArtWidth = len(line)
            }
        }

        // Move the ASCII art 20% to the right
        asciiArtX := cmdWindowWidth + (descWindowWidth-asciiArtWidth)/2 + int(0.2*float64(descWindowWidth))
        asciiArtY := (cmdWindowHeight - asciiArtHeight) / 2

        // Render the ASCII art in the background
        for y, line := range asciiArtLines {
            for x, r := range line {
                screen.SetContent(asciiArtX+x, asciiArtY+y, r, nil, tealStyle)
            }
        }

        // Displaying the commands and descriptions in the windows
        var filteredWords []KeyValuePair
        for _, word := range words {
            if strings.Contains(strings.ToLower(word.Name), strings.ToLower(inputStr)) || strings.Contains(strings.ToLower(word.Description), strings.ToLower(inputStr)) {
                filteredWords = append(filteredWords, word)
            }
        }

        // Find the closest match
        closest = findClosestMatch(inputStr, filteredWords)

        // Move the closest match to the top of the list
        if len(filteredWords) > 0 && closest.Name != "" {
            for i, word := range filteredWords {
                if word.Name == closest.Name && word.Description == closest.Description {
                    // Swap the closest match with the first element
                    filteredWords[0], filteredWords[i] = filteredWords[i], filteredWords[0]
                    break
                }
            }
        }

        // Highlight the selected command based on selectedIndex
        for i, word := range filteredWords {
            if i < cmdWindowHeight-1 {
                style := tcell.StyleDefault
                if i == selectedIndex {
                    style = highlightStyle
                }

                // Wrap the command name and description
                wrappedName := wrapText(word.Name, cmdWindowWidth-2)
                wrappedDescription := wrapText(word.Description, descWindowWidth-2)

                // Set content for wrapped command name
                currentLine := i + 1
                for _, line := range wrappedName {
                    if currentLine < cmdWindowHeight-1 {
                        for j, r := range line {
                            screen.SetContent(j+1, currentLine, r, nil, style)
                        }
                        currentLine++
                    }
                }

                // Set content for wrapped description
                currentLine = i + 1
                for _, line := range wrappedDescription {
                    if currentLine < cmdWindowHeight-1 {
                        for j, r := range line {
                            screen.SetContent(cmdWindowWidth+2+j, currentLine, r, nil, style)
                        }
                        currentLine++
                    }
                }
            }
        }

        // If the number of filtered search results is less than 10, draw a box and display flags
        if len(filteredWords) < 10 {
            // Calculate new dimensions
            newWidth := cmdWindowWidth + 78
            newHeight := cmdWindowHeight - 20
            boxXStart := cmdWindowWidth + 1
            boxXEnd := boxXStart + newWidth + 1
            boxYEnd := cmdWindowHeight
            boxYStart := boxYEnd - newHeight

            // Draw the box
            drawBorder(screen, boxXStart, boxYStart, boxXEnd, boxYEnd, tealStyle)

            // Display the flags and their descriptions for the selected command
            if selectedIndex < len(filteredWords) {
                selectedCommand := filteredWords[selectedIndex]
                flags := getFlagsForCommand(selectedCommand.Name)
                for i, flag := range flags {
                    if i < boxYEnd-boxYStart-1 {
                        y := boxYStart + 1 + i
                        screen.SetContent(boxXStart+1, y, rune(flag.Name[0]), nil, highlightStyle)
                        for j, r := range flag.Name[1:] {
                            screen.SetContent(boxXStart+2+j, y, r, nil, highlightStyle)
                        }

                        // Wrap the flag description
                        wrappedDescription := wrapText(flag.Description, newWidth-2-len(flag.Name)-1)
                        for k, line := range wrappedDescription {
                            if y+k < boxYEnd {
                                for j, r := range line {
                                    screen.SetContent(boxXStart+2+len(flag.Name)+1+j, y+k, r, nil, highlightStyle)
                                }
                            }
                        }
                    }
                }
            }
        }

        // Flushing the changes to the screen
        screen.Show()

        // Waiting for an event
        ev := screen.PollEvent()
        switch ev := ev.(type) {
        case *tcell.EventKey:
            if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
                return
            }
            if ev.Key() == tcell.KeyEnter {
                // Exit the CLI and write the command name and description to the user's command prompt
                selectedCommand := filteredWords[selectedIndex]
                screen.Fini() // Finalize the screen before printing
                fmt.Printf("%s: %s\n", selectedCommand.Name, selectedCommand.Description)
                flags := getFlagsForCommand(selectedCommand.Name)
                for _, flag := range flags {
                    fmt.Printf("  %s: %s\n", flag.Name, flag.Description)
                }
                return
            }
            if ev.Key() == tcell.KeyUp {
                if selectedIndex > 0 {
                    selectedIndex--
                }
            }
            if ev.Key() == tcell.KeyDown {
                if selectedIndex < len(filteredWords)-1 {
                    selectedIndex++
                }
            }
            if ev.Key() == tcell.KeyBackspace || ev.Key() == tcell.KeyBackspace2 {
                if len(userInput) > 0 {
                    userInput = userInput[:len(userInput)-1]
                }
            } else if ev.Rune() != 0 {
                userInput = append(userInput, ev.Rune())
            }
        case *tcell.EventError:
            fmt.Println("Error:", ev.Error())
            return
        }
    }
}