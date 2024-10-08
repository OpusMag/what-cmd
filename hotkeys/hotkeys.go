package hotkeys

// Hotkey represents a hotkey and its description
type Hotkey struct {
    Name         string
    Description string
}

// List of hotkeys with descriptions
var Words = []Hotkey{
    // Default Linux Hotkeys
    {"Linux: Ctrl + Alt + T", "Open a terminal"},
    {"Linux: Ctrl + Alt + L", "Lock the screen"},
    {"Linux: Ctrl + Alt + D", "Show desktop"},
    {"Linux: Alt + Tab", "Switch between open applications"},
    {"Linux: Alt + F2", "Run command"},
    {"Linux: Print Screen", "Take a screenshot"},
    {"Linux: Ctrl + Shift + Esc", "Open system monitor"},
    {"Linux: Ctrl + Alt + Arrow keys", "Switch workspace"},
    {"Linux: Super + L", "Lock the screen"},
    {"Linux: Super + D", "Show desktop"},
    {"Linux: Ctrl + Alt + Backspace", "Restart X server"},
    {"Linux: Ctrl + Alt + F1-F6", "Switch to virtual console"},
    {"Linux: Ctrl + Alt + F7", "Switch back to graphical session"},
    {"Linux: Alt + F4", "Close the current window"},
    {"Linux: Super + Tab", "Switch between open windows in the current workspace"},

    // Default terminal hotkeys
    {"Terminal: Ctrl + A", "Move the cursor to the beginning of the line"},
    {"Terminal: Ctrl + E", "Move the cursor to the end of the line"},
    {"Terminal: Ctrl + U", "Clear the line before the cursor"},
    {"Terminal: Ctrl + K", "Clear the line after the cursor"},
    {"Terminal: Ctrl + W", "Delete the word before the cursor"},
    {"Terminal: Ctrl + Y", "Paste the last thing to be cut (yank)"},
    {"Terminal: Alt + B", "Move the cursor backward one word"},
    {"Terminal: Alt + F", "Move the cursor forward one word"},
    {"Terminal: Ctrl + L", "Clear the screen (similar to the `clear` command)"},
    {"Terminal: Ctrl + R", "Search the command history"},
    {"Terminal: Ctrl + P", "Previous command in history"},
    {"Terminal: Ctrl + N", "Next command in history"},
    {"Terminal: !!", "Repeat the last command"},
    {"Terminal: !n", "Execute the nth command in history"},
    {"Terminal: !string", "Execute the most recent command that starts with `string`"},
    {"Terminal: Ctrl + C", "Kill the current process"},
    {"Terminal: Ctrl + Z", "Suspend the current process"},
    {"Terminal: fg", "Resume the suspended process in the foreground"},
    {"Terminal: bg", "Resume the suspended process in the background"},
    {"Terminal: Ctrl + D", "Delete the character under the cursor or exit the terminal if the line is empty"},
    {"Terminal: Ctrl + H", "Delete the character before the cursor (backspace)"},
    {"Terminal: Ctrl + T", "Swap the last two characters before the cursor"},
    {"Terminal: Alt + T", "Swap the last two words before the cursor"},
    {"Terminal: Alt + U", "Capitalize the word after the cursor"},
    {"Terminal: Alt + L", "Lowercase the word after the cursor"},
    {"Terminal: Alt + C", "Capitalize the word after the cursor"},
    {"Terminal: Tab", "Auto-complete files and directory names"},
    {"Terminal: Ctrl + _", "Undo the last action"},
    {"Terminal: Ctrl + X + E", "Open the current command in the default text editor"},

    // Kitty Terminal Hotkeys
    {"Kitty: Ctrl + Shift + Enter", "Open a new window"},
    {"Kitty: Ctrl + Shift + T", "Open a new tab"},
    {"Kitty: Ctrl + Shift + W", "Close the current tab"},
    {"Kitty: Ctrl + Shift + Left/Right", "Switch between tabs"},
    {"Kitty: Ctrl + Shift + Up/Down", "Scroll up/down"},
    {"Kitty: Ctrl + Shift + F", "Search in terminal"},
    {"Kitty: Ctrl + Shift + C", "Copy selected text"},
    {"Kitty: Ctrl + Shift + V", "Paste clipboard content"},
    {"Kitty: Ctrl + Shift + R", "Reload configuration"},
    {"Kitty: Ctrl + Shift + Q", "Quit Kitty"},
    {"Kitty: Ctrl + Shift + N", "Move to the next layout"},
    {"Kitty: Ctrl + Shift + P", "Move to the previous layout"},
    {"Kitty: Ctrl + Shift + L", "Reflow text in the current window"},
    {"Kitty: Ctrl + Shift + S", "Save the current window's scrollback buffer to a file"},

    // Tmux Terminal Hotkeys
    {"Tmux: Ctrl + B, %", "Split window vertically"},
    {"Tmux: Ctrl + B, \"", "Split window horizontally"},
    {"Tmux: Ctrl + B, Arrow keys", "Navigate between panes"},
    {"Tmux: Ctrl + B, C", "Create a new window"},
    {"Tmux: Ctrl + B, N", "Next window"},
    {"Tmux: Ctrl + B, P", "Previous window"},
    {"Tmux: Ctrl + B, W", "List all windows"},
    {"Tmux: Ctrl + B, D", "Detach session"},
    {"Tmux: Ctrl + B, R", "Reload configuration"},
    {"Tmux: Ctrl + B, X", "Kill the current pane"},
    {"Tmux: Ctrl + B, Z", "Toggle pane zoom"},
    {"Tmux: Ctrl + B, ,", "Rename the current window"},
    {"Tmux: Ctrl + B, .", "Move the current window"},
    {"Tmux: Ctrl + B, &", "Kill the current window"},
    {"Tmux: Ctrl + B, [", "Enter copy mode"},
    {"Tmux: Ctrl + B, ]", "Paste from buffer"},

    // Vim Hotkeys
    {"Vim: i", "Enter insert mode"},
    {"Vim: Esc", "Exit insert mode"},
    {"Vim: :w", "Save the file"},
    {"Vim: :q", "Quit Vim"},
    {"Vim: :wq", "Save and quit Vim"},
    {"Vim: :q!", "Quit without saving"},
    {"Vim: dd", "Delete the current line"},
    {"Vim: yy", "Yank (copy) the current line"},
    {"Vim: p", "Paste the yanked or deleted text"},
    {"Vim: u", "Undo the last action"},
    {"Vim: Ctrl + r", "Redo the undone action"},
    {"Vim: gg", "Go to the beginning of the file"},
    {"Vim: G", "Go to the end of the file"},
    {"Vim: /pattern", "Search for a pattern"},
    {"Vim: n", "Repeat the last search in the same direction"},
    {"Vim: N", "Repeat the last search in the opposite direction"},
    {"Vim: :%s/old/new/g", "Replace all occurrences of 'old' with 'new' in the file"},
    {"Vim: :e filename", "Open a file"},
    {"Vim: :bnext or :bn", "Go to the next buffer"},
    {"Vim: :bprev or :bp", "Go to the previous buffer"},
    {"Vim: :bd", "Delete a buffer"},
    {"Vim: :sp filename", "Open a file in a new split"},
    {"Vim: :vsp filename", "Open a file in a new vertical split"},
    {"Vim: Ctrl + ws", "Split the window horizontally"},
    {"Vim: Ctrl + wv", "Split the window vertically"},
    {"Vim: Ctrl + ww", "Switch between windows"},
    {"Vim: Ctrl + wq", "Quit a window"},
    {"Vim: Ctrl + wx", "Exchange current window with the next one"},
    {"Vim: :resize +N", "Increase window height by N lines"},
    {"Vim: :resize -N", "Decrease window height by N lines"},
    {"Vim: :vertical resize +N", "Increase window width by N columns"},
    {"Vim: :vertical resize -N", "Decrease window width by N columns"},
    {"Vim: :tabnew filename", "Open a file in a new tab"},
    {"Vim: gt", "Go to the next tab"},
    {"Vim: gT", "Go to the previous tab"},
    {"Vim: :tabclose or :tabc", "Close the current tab"},
    {"Vim: :tabonly", "Close all other tabs"},
    {"Vim: :tabmove N", "Move the current tab to the Nth position"},
    {"Vim: :tabedit filename", "Edit a file in a new tab"},
    {"Vim: :tabfind filename", "Open a file in a new tab if it exists"},
    {"Vim: :tabnew", "Open a new tab"},
    {"Vim: :tabnext or :tabn", "Go to the next tab"},
    {"Vim: :tabprevious or :tabp", "Go to the previous tab"},
    {"Vim: :tabfirst", "Go to the first tab"},
    {"Vim: :tablast", "Go to the last tab"},
    {"Vim: :tabdo command", "Run a command in all tabs"},
    {"Vim: :tab split", "Split the current window and open it in a new tab"},
    {"Vim: :tabclose N", "Close the Nth tab"},
    {"Vim: :tabmove +N", "Move the current tab N positions to the right"},
    {"Vim: :tabmove -N", "Move the current tab N positions to the left"},
    {"Vim: :tabnew +N", "Open a new tab and move it to the Nth position"},
    {"Vim: :tabedit +N", "Edit a file in a new tab and move it to the Nth position"},
    {"Vim: :tabfind +N", "Open a file in a new tab if it exists and move it to the Nth position"},
    {"Vim: :tabnew -N", "Open a new tab and move it to the Nth position from the end"},
    {"Vim: :tabedit -N", "Edit a file in a new tab and move it to the Nth position from the end"},
    {"Vim: :tabfind -N", "Open a file in a new tab if it exists and move it to the Nth position from the end"},
    {"Vim: :tab split +N", "Split the current window and open it in a new tab at the Nth position"},
    {"Vim: :tab split -N", "Split the current window and open it in a new tab at the Nth position from the end"},
    {"Vim: :tabclose +N", "Close the Nth tab from the current tab"},
    {"Vim: :tabclose -N", "Close the Nth tab from the last tab"},
    {"Vim: :tabonly!", "Close all other tabs except the current one"},
    {"Vim: :tabonly +N", "Close all other tabs except the Nth tab"},
    {"Vim: :tabonly -N", "Close all other tabs except the Nth tab from the last tab"},
    {"Vim: :tabmove +N", "Move the current tab to the Nth position"},
    {"Vim: :tabmove -N", "Move the current tab to the Nth position from the last tab"},

    // Nano Hotkeys
    {"Nano: Ctrl + O", "Write (save) the file"},
    {"Nano: Ctrl + X", "Exit Nano"},
    {"Nano: Ctrl + K", "Cut the current line"},
    {"Nano: Ctrl + U", "Uncut (paste) the last cut text"},
    {"Nano: Ctrl + J", "Justify the current paragraph"},
    {"Nano: Ctrl + W", "Search for a string"},
    {"Nano: Ctrl + C", "Show the current cursor position"},
    {"Nano: Ctrl + G", "Display the help text"},
    {"Nano: Ctrl + \\", "Replace a string"},
    {"Nano: Ctrl + T", "Invoke the spell checker"},
    {"Nano: Ctrl + _", "Go to a specific line number"},
    {"Nano: Alt + A", "Set a mark (start highlighting text)"},
    {"Nano: Alt + 6", "Copy the current line"},

    // Visual Studio Code Hotkeys
    {"VSCode: Ctrl + P", "Quick open a file"},
    {"VSCode: Ctrl + Shift + P", "Open the command palette"},
    {"VSCode: Ctrl + ,", "Open settings"},
    {"VSCode: Ctrl + B", "Toggle the sidebar"},
    {"VSCode: Ctrl + `", "Toggle the integrated terminal"},
    {"VSCode: Ctrl + K Ctrl + S", "Open keyboard shortcuts"},
    {"VSCode: Ctrl + Shift + N", "New window"},
    {"VSCode: Ctrl + W", "Close the current editor"},
    {"VSCode: Ctrl + Shift + T", "Reopen the last closed editor"},
    {"VSCode: Ctrl + Tab", "Switch between open editors"},
    {"VSCode: Ctrl + /", "Toggle line comment"},
    {"VSCode: Ctrl + Shift + A", "Toggle block comment"},
    {"VSCode: Alt + Up/Down", "Move line up/down"},
    {"VSCode: Shift + Alt + Up/Down", "Copy line up/down"},
    {"VSCode: Ctrl + Shift + K", "Delete line"},
    {"VSCode: Ctrl + Enter", "Insert line below"},
    {"VSCode: Ctrl + Shift + Enter", "Insert line above"},
    {"VSCode: Ctrl + D", "Add selection to next find match"},
    {"VSCode: Ctrl + L", "Select current line"},
    {"VSCode: Ctrl + Shift + L", "Select all occurrences of current selection"},
    {"VSCode: F2", "Rename symbol"},
    {"VSCode: Ctrl + Shift + O", "Go to symbol"},
    {"VSCode: Ctrl + G", "Go to line"},
    {"VSCode: F12", "Go to definition"},
    {"VSCode: Alt + F12", "Peek definition"},
    {"VSCode: Ctrl + K Ctrl + 0", "Fold all regions"},
    {"VSCode: Ctrl + K Ctrl + J", "Unfold all regions"},
    {"VSCode: Ctrl + Shift + M", "Show problems"},
    {"VSCode: F8", "Go to next error or warning"},
    {"VSCode: Shift + F8", "Go to previous error or warning"},
    {"VSCode: Ctrl + Shift + B", "Run build task"},
    {"VSCode: Ctrl + Shift + U", "Show output panel"},
    {"VSCode: Ctrl + Shift + Y", "Show debug console"},
    {"VSCode: F5", "Start debugging"},
    {"VSCode: Shift + F5", "Stop debugging"},
    {"VSCode: F9", "Toggle breakpoint"},
    {"VSCode: F10", "Step over"},
    {"VSCode: F11", "Step into"},
    {"VSCode: Shift + F11", "Step out"},

    // i3 Window Manager Hotkeys
    {"i3: Mod + Enter", "Open a terminal"},
    {"i3: Mod + D", "Open dmenu"},
    {"i3: Mod + J/K/L/; or Arrow keys", "Navigate between windows"},
    {"i3: Mod + H/V", "Split horizontally/vertically"},
    {"i3: Mod + F", "Toggle fullscreen"},
    {"i3: Mod + Shift + Q", "Kill focused window"},
    {"i3: Mod + Shift + R", "Reload i3 configuration"},
    {"i3: Mod + Shift + E", "Exit i3"},
    {"i3: Mod + 1-9", "Switch to workspace 1-9"},
    {"i3: Mod + Shift + 1-9", "Move focused window to workspace 1-9"},
    {"i3: Mod + S", "Toggle stacking layout"},
    {"i3: Mod + W", "Toggle tabbed layout"},
    {"i3: Mod + E", "Toggle split layout"},
    {"i3: Mod + Shift + Space", "Toggle floating mode"},
    {"i3: Mod + Space", "Focus the parent container"},
    {"i3: Mod + Shift + C", "Reload i3 configuration"},

    // Awesome Window Manager Hotkeys
    {"Awesome: Mod4 + Enter", "Open a terminal"},
    {"Awesome: Mod4 + R", "Run prompt"},
    {"Awesome: Mod4 + J/K", "Focus next/previous window"},
    {"Awesome: Mod4 + H/L", "Resize window"},
    {"Awesome: Mod4 + F", "Toggle fullscreen"},
    {"Awesome: Mod4 + Shift + C", "Close focused window"},
    {"Awesome: Mod4 + Control + R", "Reload Awesome"},
    {"Awesome: Mod4 + Control + Q", "Quit Awesome"},
    {"Awesome: Mod4 + 1-9", "Switch to tag 1-9"},
    {"Awesome: Mod4 + Shift + 1-9", "Move focused window to tag 1-9"},
    {"Awesome: Mod4 + W", "Show main menu"},
    {"Awesome: Mod4 + M", "Maximize window"},
    {"Awesome: Mod4 + N", "Minimize window"},
    {"Awesome: Mod4 + P", "Show program launcher"},
    {"Awesome: Mod4 + S", "Take a screenshot"},
    {"Awesome: Mod4 + Tab", "Switch between clients"},

    // Windows Terminal Hotkeys
    {"Windows Terminal: Ctrl + Shift + 1", "Open a new tab with the default profile"},
    {"Windows Terminal: Ctrl + Shift + 2", "Open a new tab with the second profile"},
    {"Windows Terminal: Ctrl + Shift + W", "Close the current tab"},
    {"Windows Terminal: Ctrl + Shift + D", "Duplicate the current tab"},
    {"Windows Terminal: Ctrl + Shift + T", "Open a new tab"},
    {"Windows Terminal: Ctrl + Shift + P", "Open the command palette"},
    {"Windows Terminal: Ctrl + Shift + F", "Find text in the terminal"},
    {"Windows Terminal: Ctrl + Shift + Plus", "Increase font size"},
    {"Windows Terminal: Ctrl + Shift + Minus", "Decrease font size"},
    {"Windows Terminal: Alt + Shift + D", "Split the current pane horizontally"},
    {"Windows Terminal: Alt + Shift + Plus", "Split the current pane vertically"},
    {"Windows Terminal: Alt + Arrow Keys", "Move focus between panes"},
    {"Windows Terminal: Ctrl + Shift + Space", "Open the context menu"},
    {"Windows Terminal: Ctrl + Shift + C", "Copy selected text"},
    {"Windows Terminal: Ctrl + Shift + V", "Paste from clipboard"},

    // PowerShell Hotkeys
    {"PowerShell: Ctrl + C", "Interrupt the current command"},
    {"PowerShell: Ctrl + D", "Exit the shell"},
    {"PowerShell: Ctrl + L", "Clear the screen"},
    {"PowerShell: Ctrl + R", "Search command history"},
    {"PowerShell: Ctrl + A", "Select all text"},
    {"PowerShell: Ctrl + E", "Move to the end of the line"},
    {"PowerShell: Ctrl + U", "Delete from cursor to the beginning of the line"},
    {"PowerShell: Ctrl + K", "Delete from cursor to the end of the line"},
    {"PowerShell: Ctrl + W", "Delete the word before the cursor"},
    {"PowerShell: Alt + F", "Move cursor forward one word"},
    {"PowerShell: Alt + B", "Move cursor backward one word"},
    {"PowerShell: Alt + D", "Delete the word after the cursor"},
    {"PowerShell: Shift + Insert", "Paste from clipboard"},
    {"PowerShell: Shift + Up Arrow", "Scroll up through command history"},
    {"PowerShell: Shift + Down Arrow", "Scroll down through command history"},

    // Postman Hotkeys
    {"Postman: Ctrl + N", "New tab"},
    {"Postman: Ctrl + T", "New request tab"},
    {"Postman: Ctrl + W", "Close tab"},
    {"Postman: Ctrl + Shift + N", "New window"},
    {"Postman: Ctrl + Shift + W", "Close window"},
    {"Postman: Ctrl + Enter", "Send request"},
    {"Postman: Ctrl + S", "Save request"},
    {"Postman: Ctrl + F", "Find"},
    {"Postman: Ctrl + Shift + F", "Find and replace"},
    {"Postman: Ctrl + Shift + C", "Toggle comment"},
    {"Postman: Ctrl + Shift + M", "Toggle sidebar"},

    // macOS Hotkeys
    {"macOS: Command + C", "Copy"},
    {"macOS: Command + V", "Paste"},
    {"macOS: Command + X", "Cut"},
    {"macOS: Command + Z", "Undo"},
    {"macOS: Command + Shift + Z", "Redo"},
    {"macOS: Command + A", "Select all"},
    {"macOS: Command + F", "Find"},
    {"macOS: Command + H", "Hide application"},
    {"macOS: Command + Q", "Quit application"},
    {"macOS: Command + Space", "Spotlight search"},
    {"macOS: Command + Tab", "Switch applications"},
    {"macOS: Command + Option + Esc", "Force quit applications"},

    // macOS Terminal Hotkeys
    {"macOS Terminal: Command + T", "New tab"},
    {"macOS Terminal: Command + N", "New window"},
    {"macOS Terminal: Command + W", "Close tab"},
    {"macOS Terminal: Command + Shift + W", "Close window"},
    {"macOS Terminal: Command + K", "Clear screen"},
    {"macOS Terminal: Command + Arrow Up", "Scroll up"},
    {"macOS Terminal: Command + Arrow Down", "Scroll down"},
    {"macOS Terminal: Command + Arrow Left", "Move cursor to beginning of line"},
    {"macOS Terminal: Command + Arrow Right", "Move cursor to end of line"},
    {"macOS Terminal: Option + Arrow Left", "Move cursor one word left"},
    {"macOS Terminal: Option + Arrow Right", "Move cursor one word right"},
    {"macOS Terminal: Control + C", "Interrupt current command"},
    {"macOS Terminal: Control + D", "Exit the shell"},
    {"macOS Terminal: Control + L", "Clear the screen"},
    {"macOS Terminal: Control + R", "Search command history"},
    {"macOS Terminal: Control + A", "Move to the beginning of the line"},
    {"macOS Terminal: Control + E", "Move to the end of the line"},
    {"macOS Terminal: Control + U", "Delete from cursor to the beginning of the line"},
    {"macOS Terminal: Control + K", "Delete from cursor to the end of the line"},
    {"macOS Terminal: Control + W", "Delete the word before the cursor"},
    {"macOS Terminal: Control + Y", "Paste the last cut text"},
    {"macOS Terminal: Control + T", "Transpose characters"},

    // zsh Hotkeys
    {"zsh: Control + A", "Move to the beginning of the line"},
    {"zsh: Control + E", "Move to the end of the line"},
    {"zsh: Control + U", "Delete from cursor to the beginning of the line"},
    {"zsh: Control + K", "Delete from cursor to the end of the line"},
    {"zsh: Control + W", "Delete the word before the cursor"},
    {"zsh: Control + Y", "Paste the last cut text"},
    {"zsh: Control + T", "Transpose characters"},
    {"zsh: Control + L", "Clear the screen"},
    {"zsh: Control + R", "Search command history"},
    {"zsh: Control + C", "Interrupt current command"},
    {"zsh: Control + D", "Exit the shell"},
    {"zsh: Alt + F", "Move cursor forward one word"},
    {"zsh: Alt + B", "Move cursor backward one word"},
    {"zsh: Alt + D", "Delete the word after the cursor"},
    {"zsh: Alt + Backspace", "Delete the word before the cursor"},
    {"zsh: Control + X Control + E", "Open the current command in the default editor"},
    {"zsh: Control + P", "Previous command in history"},
    {"zsh: Control + N", "Next command in history"},
    {"zsh: Control + S", "Search forward in command history"},
    {"zsh: Control + Q", "Resume output after Control + S"},
    {"zsh: Control + Z", "Suspend the current foreground job"},
    {"zsh: Control + _", "Undo the last editing command"},
}