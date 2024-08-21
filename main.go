package main

import (
    "flag"
    "fmt"
    "os"
    "strings"
    "github.com/gdamore/tcell/v2"
    "what-cmd/commands"
    "what-cmd/flags"
)

// KeyValuePair structure
type KeyValuePair struct {
    Name        string
    Description string
}

// Convert flags.Flag to KeyValuePair
func convertFlagsToKeyValuePairs(flags []flags.Flag) []KeyValuePair {
    var keyValuePairs []KeyValuePair
    for _, f := range flags {
        keyValuePairs = append(keyValuePairs, KeyValuePair{Name: f.Name, Description: f.Description})
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

// Finds the minimum of two integers
func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func main() {
    // Setting the CLI name
    flag.CommandLine.Usage = func() {
        fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
        flag.PrintDefaults()
    }

    // Defining flags
    input := flag.String("input", "", "input word to search for")
    useFlags := flag.Bool("flags", false, "search in flags instead of commands")
    flag.Parse()

    var words []KeyValuePair
    if *useFlags {
        words = convertFlagsToKeyValuePairs(flags.Words)
    } else {
        words = convertCommandsToKeyValuePairs(commands.Words)
    }

    if *input != "" {
        closest := findClosestMatch(*input, words)
        fmt.Printf("Closest match: %s - %s\n", closest.Name, closest.Description)
        return
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

    // Main loop
    for {
        // Clearing the screen
        screen.Clear()

        // Displaying the prompt
        prompt := "Enter a command to search for (type 'exit' to quit): "
        for i, r := range prompt {
            screen.SetContent(i, 0, r, nil, tcell.StyleDefault)
        }

        // Displaying the user input
        for i, r := range userInput {
            screen.SetContent(i, 1, r, nil, tcell.StyleDefault)
        }

        // Finding and displaying the closest match
        inputStr := string(userInput)
        closest := findClosestMatch(inputStr, words)
        matchStr := fmt.Sprintf(
            "Closest match: %s - %s",
            closest.Name,
            closest.Description,
        )
        for i, r := range matchStr {
            screen.SetContent(i, 2, r, nil, tcell.StyleDefault)
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
                screen.Fini()
                fmt.Printf("\n%s - %s\n%s ", closest.Name, closest.Description, closest.Name)
                return
            }
            if ev.Key() == tcell.KeyBackspace || ev.Key() == tcell.KeyBackspace2 {
                if len(userInput) > 0 {
                    userInput = userInput[:len(userInput)-1]
                }
            } else {
                userInput = append(userInput, ev.Rune())
            }
        case *tcell.EventError:
            fmt.Println("Error:", ev.Error())
            return
        }
    }
}