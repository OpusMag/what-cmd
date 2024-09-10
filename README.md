# what-cmd
What-cmd is a cross-platform (windows, linux and macOS) command line tool written in Go that gives you an easy way to find the terminal command (and its flags) you're looking for when you've forgotten it or if you're new to the terminal. It is inspired by which-key and lazygit.

![what-cmd-demo](https://github.com/user-attachments/assets/4261dfce-455f-4f73-8115-57b4c07a32d2)

Future releases: In future releases the command list may be expanded further. Improvements to the UX/UI may be considered as well as further functionality.

To use this tool:

Install dependencies first (Git and Go programming language).  

Linux:
Depending on your package manager:
1. **Git:**
   1. sudo pacman -S git -y
   2. sudo dnf install git-all
   3. sudo apt install git-all

2. **Git:**
   1. sudo pacman -S go
   2. sudo dnf install go
   3. sudo apt install go

Windows:
1. **Git:**
   1. go to https://git-scm.com/download/win and the download should start automatically. Then install it.

2. **Git:**
   1. go to https://golang.org/dl/ and download the latest version of go. 
   2. Follow this guide to install properly: https://www.geeksforgeeks.org/how-to-install-go-on-windows/

macOS
1. **Git:**
   - Using Homebrew (Recommended):
     ```sh
     brew install git
     ```
   - Using Xcode Command Line Tools:
     ```sh
     xcode-select --install
     ```

2. **Go:**
   - Using Homebrew (Recommended):
     ```sh
     brew install go
     ```
   - Manual Installation:
     - Download the latest version of Go from [https://golang.org/dl/](https://golang.org/dl/)
     - Open the downloaded `.pkg` file and follow the instructions to install Go.

Linux and macOS instructions for installing what-cmd:

Option A: 
   1. Go to the latest release at https://github.com/OpusMag/what-cmd and download the binary. 
   2. Eun chmod +x ~/Downloads/what-cmd (or whatever dir you downloaded it to) 
   3. Move it to /usr/local/bin/ by running sudo mv what-cmd /usr/local/bin/

Option B: 
   1. Enter your terminal. 
   2. Git clone this repository either by using the URL or SSH (you can find this by clicking the green button in the upper right that says '<>Code'). 
   3. Then cd (change directory) into the repository. 
   4. Run the following commands to build and run the CLI:

      go mod tidy
      go build -o what-cmd main.go
      ./what-cmd

      If you want to use it without the ./ prefix you can do the following:
      while in the what-cmd directory and after building the binary:
      sudo mv what-cmd /usr/local/bin/

      Confirm it has worked:
      what-cmd --help (this also shows the available flags)

Windows instructions:

Option A: 
   1. Go to the latest release at https://github.com/OpusMag/what-cmd, download the binary what-cmd.exe.
   2. Move the binary to a directory, for example C:\Tools.
   3. Add the directory to PATH by pressing Win + X and selecting System, then click on Advanced system settings. 
   4. In the System properties window, click on the Environment variables button. 
   5. In the Environment Variables window, find the Path variable in the System variables section and select it and click Edit. 
   6. In the Edit Environment Variable window, click New and add the path to the directory where you placed what-cmd.exe (for example c:\Tools), then click OK to close. 
   7. Verify it's working by opening a terminal and running 'what-cmd --help'. 

Option B: 
   1. 'git clone https://github.com/OpusMag/what-cmd' then cd into the repository. 
   2. Run 'go mod tidy' in the terminal to make sure dependencies are installed. 
   3. Then run 'go build -o what-cmd.exe main.go'. Then run the tool with '.\what-cmd.exe'

General use of the tool(both windows and linux):

1. Options for the command (run "what-cmd -hotkeys" for example to search hotkeys for the terminal)

   -flags (flags for the commands. NB! Flags will now appear when you search for commands just by running what-cmd without flags)
   -hotkeys (linux terminal hotkeys)

2. By running ./what-cmd (or what-cmd) you'll now be able to search for commands either by name or description. If you append a flag (e.g. what-cmd -flags) you'll be able to search for the flags that are valid for each command.

3. When you've found the command (or flag or hotkey), hit enter and you'll be sent back to your terminal prompt and the command you found along with its description will be printed to your terminal.
