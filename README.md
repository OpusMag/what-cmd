# what-cmd
What-cmd is a command line tool written in Golang that gives you an easy way to find the linux terminal command you're looking for when you've forgotten it or if you're new.

Current version: v0.1 released 20.08.2024

Future releases: In future releases the command list will be expanded. In addition common flags for commands may be added so you can find the right flags and a description for them. The GUI may also be improved upon further.

To use this tool:

1. Enter your terminal

2. Git clone this repository either by using the URL or SSH (you can find this by clicking the green button in the upper right that says '<>Code').

3. Then cd (change directory) into the repository

4. Run the following commands to build and run the CLI:

   go mod tidy
   go build -o what-cmd main.go
   ./what-cmd

7. You can now type in a description of the command or if you remember part of the command name, you can try that and see if the description matches.

8. When you've found the command, hit enter and you'll be sent back to your terminal prompt and the command you found along with its description will be printed to your terminal.
