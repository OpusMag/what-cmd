# what-cmd
What-cmd is a command line tool written in Go that gives you an easy way to find the linux terminal command (and its flags) you're looking for when you've forgotten it or if you're new to the linux terminal.

Future releases: In future releases the command list may be expanded further. In addition, improvements to the UX/UI may be considered as well as further functionality.

To use this tool:

Option A: Go to the latest release, download the binary, run chmod +x ~/Downloads/what-cmd (or whatever dir you downloaded it to) and then move it to /usr/local/bin/ by running sudo mv what-cmd /usr/local/bin/

Option B:

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

5. Options for the command (run "what-cmd -hotkeys" for example to search hotkeys for the terminal)

   -flags (flags for the commands. NB! Flags will now appear when you search for commands just by running what-cmd without flags)
   -hotkeys (linux terminal hotkeys)

6. By running ./what-cmd (or what-cmd) you'll now be able to search for commands either by name or description. If you append a flag (e.g. what-cmd -flags) you'll be able to search for the flags that are valid for each command.

7. When you've found the command (or flag or hotkey), hit enter and you'll be sent back to your terminal prompt and the command you found along with its description will be printed to your terminal.
