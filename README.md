# what-cmd
What-cmd is a cross-platform (works on windows and linux) command line tool written in Go that gives you an easy way to find the terminal command (and its flags) you're looking for when you've forgotten it or if you're new to the terminal. It is inspired by which-key and lazygit.

Future releases: In future releases the command list may be expanded further. In addition, improvements to the UX/UI may be considered as well as further functionality.

To use this tool:

Install dependencies first (Git and Go programming language).  

in your terminal (linux):

Depending on your package manager:
Git:
sudo pacman -S git -y
sudo dnf install git-all
sudo apt install git-all

Go:
sudo pacman -S go
sudo dnf install go
sudo apt install go

Windows:
Git:
go to https://git-scm.com/download/win and the download should start automatically. Then install it.

Go (in the terminal):
go to https://golang.org/dl/ and download the latest version of go. Follow this guide to install properly: https://www.geeksforgeeks.org/how-to-install-go-on-windows/

Linux instructions:

Option A: Go to the latest release at https://github.com/OpusMag/what-cmd , download the binary, run chmod +x ~/Downloads/what-cmd (or whatever dir you downloaded it to) and then move it to /usr/local/bin/ by running sudo mv what-cmd /usr/local/bin/

Option B: Enter your terminal. Git clone this repository either by using the URL or SSH (you can find this by clicking the green button in the upper right that says '<>Code'). Then cd (change directory) into the repository. Run the following commands to build and run the CLI:

   go mod tidy
   go build -o what-cmd main.go
   ./what-cmd

   If you want to use it without the ./ prefix you can do the following:
   while in the what-cmd directory and after building the binary:
   sudo mv what-cmd /usr/local/bin/

   Confirm it has worked:
   what-cmd --help (this also shows the available flags)

Windows instructions:

Option A: Go to the latest release at https://github.com/OpusMag/what-cmd, download the binary what-cmd.exe. Move the binary to a directory, for example C:\Tools. Add the directory to PATH by pressing Win + X and selecting System, then click on Advanced system settings. In the System properties window, click on the Environment variables button. In the Environment Variables window, find the Path variable in the System variables section and select it. Click Edit. In the Edit Environment Variable window, click New and add the path to the directory where you placed what-cmd.exe (for example c:\Tools), then click OK to close. Verify it's working by opening a terminal and running what-cmd --help. 

Option B: git clone https://github.com/OpusMag/what-cmd then cd into the repository. Run go mod tidy in the terminal to make sure dependencies are installed. Then run go build -o what-cmd.exe main.go. Then run the tool with .\what-cmd.exe

Further use of the tool(both windows and linux):

1. Options for the command (run "what-cmd -hotkeys" for example to search hotkeys for the terminal)

   -flags (flags for the commands. NB! Flags will now appear when you search for commands just by running what-cmd without flags)
   -hotkeys (linux terminal hotkeys)

2. By running ./what-cmd (or what-cmd) you'll now be able to search for commands either by name or description. If you append a flag (e.g. what-cmd -flags) you'll be able to search for the flags that are valid for each command.

3. When you've found the command (or flag or hotkey), hit enter and you'll be sent back to your terminal prompt and the command you found along with its description will be printed to your terminal.
