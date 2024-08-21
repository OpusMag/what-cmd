package hotkeys

// Hotkey represents a hotkey and its description
type Hotkey struct {
    Name         string
    Description string
}

// List of hotkeys with descriptions
var Words = []Hotkey{
    {"Ctrl + A", "Move the cursor to the beginning of the line"},
    {"Ctrl + E", "Move the cursor to the end of the line"},
    {"Ctrl + U", "Clear the line before the cursor"},
    {"Ctrl + K", "Clear the line after the cursor"},
    {"Ctrl + W", "Delete the word before the cursor"},
    {"Ctrl + Y", "Paste the last thing to be cut (yank)"},
    {"Alt + B", "Move the cursor backward one word"},
    {"Alt + F", "Move the cursor forward one word"},
    {"Ctrl + L", "Clear the screen (similar to the `clear` command)"},
    {"Ctrl + R", "Search the command history"},
    {"Ctrl + P", "Previous command in history"},
    {"Ctrl + N", "Next command in history"},
    {"!!", "Repeat the last command"},
    {"!n", "Execute the nth command in history"},
    {"!string", "Execute the most recent command that starts with `string`"},
    {"Ctrl + C", "Kill the current process"},
    {"Ctrl + Z", "Suspend the current process"},
    {"fg", "Resume the suspended process in the foreground"},
    {"bg", "Resume the suspended process in the background"},
    {"Ctrl + D", "Delete the character under the cursor or exit the terminal if the line is empty"},
    {"Ctrl + H", "Delete the character before the cursor (backspace)"},
    {"Ctrl + T", "Swap the last two characters before the cursor"},
    {"Alt + T", "Swap the last two words before the cursor"},
    {"Alt + U", "Capitalize the word after the cursor"},
    {"Alt + L", "Lowercase the word after the cursor"},
    {"Alt + C", "Capitalize the word after the cursor"},
    {"Tab", "Auto-complete files and directory names"},
    {"Ctrl + _", "Undo the last action"},
    {"Ctrl + X + E", "Open the current command in the default text editor"},
}