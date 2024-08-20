package main

import (
    "flag"
    "fmt"
    "os"
    "github.com/gdamore/tcell/v2"
)

// A command and its description
type Command struct {
    Name        string
    Description string
}

// List of commands with descriptions
var words = []Command{
    {"sudo", "Super User Do. Allows a user to run commands as an administrator"},
    {"cat", "Concatenate and display the content of files"},
    {"ls", "List directory contents"},
    {"cd", "Change the current directory"},
    {"pwd", "Print the current working directory"},
    {"touch", "Create an empty file"},
    {"mkdir", "Create a new directory"},
    {"rm", "Remove files or directories"},
    {"cp", "Copy files or directories"},
    {"mv", "Move or rename files or directories"},
    {"clear", "Clear the terminal screen"},
    {"exit", "Exit the shell"},
    {"echo", "Display a line of text"},
    {"man", "Display the manual for a command"},
    {"grep", "Search text using patterns"},
    {"find", "Search for files in a directory hierarchy"},
    {"chmod", "Change file modes or Access Control Lists"},
    {"chown", "Change file owner and group"},
    {"df", "Report file system disk space usage"},
    {"du", "Estimate file space usage"},
    {"head", "Output the first part of files"},
    {"tail", "Output the last part of files"},
    {"tar", "Store, list or extract files in an archive"},
    {"wget", "Retrieve files from the web"},
    {"curl", "Transfer data from or to a server"},
    {"ps", "Report a snapshot of current processes"},
    {"kill", "Send a signal to a process"},
    {"top", "Display Linux tasks"},
    {"htop", "Interactive process viewer"},
    {"ssh", "OpenSSH SSH client (remote login program)"},
    {"scp", "Secure copy (remote file copy program)"},
    {"nano", "Simple text editor"},
    {"vim", "Vi IMproved, a programmer's text editor"},
    {"apt-get", "A package handling utility for Debian-based distributions"},
    {"yum", "Package manager for RPM-based distributions"},
    {"ping", "Send ICMP ECHO_REQUEST to network hosts"},
    {"ifconfig", "Configure a network interface"},
    {"ip", "Show/manipulate routing, devices, policy routing and tunnels"},
    {"netstat", "Print network connections, routing tables, interface statistics, masquerade connections, and multicast memberships"},
    {"ss", "Another utility to investigate sockets"},
    {"df", "Report file system disk space usage"},
    {"mount", "Mount a file system"},
    {"umount", "Unmount file systems"},
    {"uname", "Print system information"},
    {"uptime", "Tell how long the system has been running"},
    {"whoami", "Print the current user id and name"},
    {"hostname", "Show or set the system's host name"},
    {"history", "Show the command history"},
    {"alias", "Create an alias for a command"},
    {"unalias", "Remove an alias"},
    {"env", "Show the environment variables"},
    {"export", "Set an environment variable"},
    {"source", "Execute commands from a file in the current shell"},
    {"sudo", "Execute a command as another user"},
    {"su", "Switch user"},
    {"passwd", "Change user password"},
    {"useradd", "Create a new user"},
    {"usermod", "Modify a user account"},
    {"userdel", "Delete a user account"},
    {"groupadd", "Create a new group"},
    {"groupmod", "Modify a group"},
    {"groupdel", "Delete a group"},
    {"crontab", "Schedule periodic background jobs"},
    {"at", "Schedule a command to run once at a later time"},
    {"jobs", "List active jobs"},
    {"bg", "Place a job in the background"},
    {"fg", "Place a job in the foreground"},
    {"killall", "Kill processes by name"},
    {"xargs", "Build and execute command lines from standard input"},
    {"tee", "Read from standard input and write to standard output and files"},
    {"diff", "Compare files line by line"},
    {"patch", "Apply a diff file to an original"},
    {"sed", "Stream editor for filtering and transforming text"},
    {"awk", "Pattern scanning and processing language"},
    {"tr", "Translate or delete characters"},
    {"cut", "Remove sections from each line of files"},
    {"sort", "Sort lines of text files"},
    {"uniq", "Report or omit repeated lines"},
    {"wc", "Print newline, word, and byte counts for each file"},
    {"basename", "Strip directory and suffix from filenames"},
    {"dirname", "Strip last component from file name"},
    {"readlink", "Print value of a symbolic link or canonical file name"},
    {"ln", "Make links between files"},
    {"stat", "Display file or file system status"},
    {"file", "Determine file type"},
    {"strings", "Print the strings of printable characters in files"},
    {"md5sum", "Compute and check MD5 message digest"},
    {"sha256sum", "Compute and check SHA256 message digest"},
    {"gzip", "Compress files"},
    {"gunzip", "Decompress files"},
    {"bzip2", "Compress files with bzip2"},
    {"bunzip2", "Decompress files with bzip2"},
    {"zip", "Package and compress files"},
    {"unzip", "Extract compressed files"},
    {"rar", "Archive manager for .rar files"},
    {"unrar", "Extract .rar files"},
    {"7z", "File archiver with high compression ratio"},
    {"7za", "Standalone version of 7z"},
    {"7zr", "Reduced version of 7z"},
}

// Find the closest match to the input
func findClosestMatch(input string) Command {
    var closest Command
    var minDistance = -1

    for _, word := range words {
        distance := distanceAtoB(input, word.Name)
        if minDistance == -1 || distance < minDistance {
            minDistance = distance
            closest = word
        }
    }

    return closest
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

func main() {
  // Setting the CLI name
  flag.CommandLine.Usage = func() {
      fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
      flag.PrintDefaults()
  }

  // Defining flags
  input := flag.String("input", "", "input word to search for")
  flag.Parse()

  if *input != "" {
      closest := findClosestMatch(*input)
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
      closest := findClosestMatch(inputStr)
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