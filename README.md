# what-cmd
What-cmd is a command line tool written in Golang that gives you an easy way to find the linux terminal command you're looking for when you've forgotten it or if you're new.

Current version: v0.2 released 21.08.2024

v0.2 improvements: Improved search algorithm to only search for exact matches for commands and otherwise prioritize matches with the description. Also added -flags option (read usage info below) and expanded command list.

Future releases: In future releases the command list may be expanded further. In addition improvements to the aesthetic may be considered as well as further functionality.

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
   what-cmd --help

5. You can now type in a description of the command or if you remember part of the command name, you can try that and see if the description matches.

6. When you've found the command, hit enter and you'll be sent back to your terminal prompt and the command you found along with its description will be printed to your terminal.

7. You can now run './what-cmd -flags' to find flags for your commands. If you don't remember what flags there are for your command you can run this command to check. Use ./what-cmd first to find the command you're looking for if you don't remember the name.
