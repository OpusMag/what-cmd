# what-cmd
What-cmd is a command line tool written in Golang that gives you an easy way to find the linux terminal command you're looking for when you've forgotten it or if you're new.

Current version: v0.2 released 21.08.2024

v0.2 improvements: Improved search algorithm to only search for exact matches for commands and otherwise prioritize matches with the description. Also added -flags, -git and -hotkeys option (read usage info below) and expanded command list.

Future releases: In future releases the command list may be expanded further. In addition, improvements to the aesthetic may be considered as well as further functionality.

To use this tool:

1. Enter your terminal

2. Git clone this repository either by using the URL or SSH (you can find this by clicking the green button in the upper right that says '<>Code').

3. Then cd (change directory) into the repository

4. Run the following commands to build and run the CLI:

   go mod tidy
   go build -o what-cmd main.go
   ./what-cmd

   If you want to use it without the ./ prefix you can do the following:
   while in the what-cmd directory and after building the binary:
   sudo mv what-cmd /usr/local/bin/

   Confirm it has worked:
   what-cmd --help (this also shows the available flags)

   -flags (flags for common linux terminal commands)
   -git (common git commands)
   -hotkeys (linux terminal hotkeys)

5. By running ./what-cmd (or what-cmd) you'll now be able to search for commands either by name or description. If you append a flag (e.g. what-cmd -flags) you'll be able to search for the flags that are valid for each command.

6. When you've found the command (or flag or hotkey), hit enter and you'll be sent back to your terminal prompt and the command you found along with its description will be printed to your terminal.
