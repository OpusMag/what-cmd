# what-cmd

What-cmd is a cross-platform (Windows, Linux, and macOS) command-line tool written in Go that provides an intuitive way to discover and search for terminal commands, flags, and hotkeys. Whether you've forgotten a command or you're new to the terminal, what-cmd makes it easy to find what you need. The tool is inspired by which-key and lazygit, featuring both built-in knowledge and dynamic system discovery.

![what-cmd-demo](https://github.com/user-attachments/assets/4261dfce-455f-4f73-8115-57b4c07a32d2)

## Features

### **Multi-Mode Search**
- **Commands**: Search through common terminal commands
- **Flags**: Find command-line flags and options
- **Hotkeys**: Discover keyboard shortcuts for various applications

### **System Discovery**
- **Automatic Detection**: Scans your system PATH for installed CLI tools
- **Shell Integration**: Discovers your custom aliases and functions from shell configuration files
- **Safe by Default**: Runs in safe mode with comprehensive security controls
- **Configurable**: Customize scan locations and behavior through configuration files

### **User Interface**
- **Interactive Terminal UI**: Clean, organized display with real-time search
- **Smart Ranking**: Intelligent scoring system for relevant results
- **Grouped Results**: Organize findings by source (built-in, system, user-defined)
- **Context-Aware**: Shows relevant flags for selected commands

## Installation

### Prerequisites

Install **Git** and **Go** first:

#### Linux
Depending on your package manager:

**Git:**
```bash
# Arch/Manjaro
sudo pacman -S git

# Fedora/RHEL/CentOS
sudo dnf install git-all

# Ubuntu/Debian
sudo apt install git-all
```

**Go:**
```bash
# Arch/Manjaro
sudo pacman -S go

# Fedora/RHEL/CentOS
sudo dnf install go

# Ubuntu/Debian
sudo apt install golang-go
```

#### Windows

**Git:**
1. Visit [https://git-scm.com/download/win](https://git-scm.com/download/win)
2. Download and install the latest version

**Go:**
1. Visit [https://golang.org/dl/](https://golang.org/dl/)
2. Download the latest Windows installer
3. Follow the [installation guide](https://www.geeksforgeeks.org/how-to-install-go-on-windows/)

#### macOS

**Git:**
```bash
# Using Homebrew (Recommended)
brew install git

# Or using Xcode Command Line Tools
xcode-select --install
```

**Go:**
```bash
# Using Homebrew (Recommended)
brew install go

# Or download from https://golang.org/dl/
```

### Install what-cmd

#### Option A: Download Binary (Recommended)

**Linux/macOS:**
1. Download the latest binary from [GitHub Releases](https://github.com/OpusMag/what-cmd/releases)
2. Make it executable: `chmod +x ~/Downloads/what-cmd`
3. Move to system path: `sudo mv ~/Downloads/what-cmd /usr/local/bin/`
4. Verify installation: `what-cmd --help`

**Windows:**
1. Download `what-cmd.exe` from [GitHub Releases](https://github.com/OpusMag/what-cmd/releases)
2. Create a directory: `C:\Tools`
3. Move the binary to `C:\Tools\what-cmd.exe`
4. Add `C:\Tools` to your PATH environment variable:
   - Press `Win + X` → System → Advanced system settings
   - Environment Variables → System variables → Path → Edit → New
   - Add `C:\Tools` → OK
5. Verify installation: `what-cmd --help`

#### Option B: Build from Source

```bash
# Clone the repository
git clone https://github.com/OpusMag/what-cmd.git
cd what-cmd

# Install dependencies
go mod tidy

# Build the application
go build -o what-cmd main.go  # Linux/macOS
go build -o what-cmd.exe main.go  # Windows

# Run locally
./what-cmd  # Linux/macOS
.\what-cmd.exe  # Windows

# Install globally (Linux/macOS)
sudo mv what-cmd /usr/local/bin/
```

## Usage

### Basic Commands

```bash
# Default mode - search built-in commands (SAFE)
what-cmd

# Search for flags
what-cmd -flags

# Search for hotkeys
what-cmd -hotkeys

# Use custom configuration
what-cmd -config=/path/to/config.json
```

## Security & Safety

When designing what-cmd, security was the top priority. However, it's crucial to understand that adding system discovery was complex and initially caused an infinite loop with cascading resource consumption on Windows. Therefore, it's **VERY IMPORTANT** that you are careful with system discovery for this tool. If you don't need system discovery, run what-cmd in safe mode.

**What-cmd offers four distinct security modes:**

### **Safe Mode** (`what-cmd` or `what-cmd -safe`) - **DEFAULT**
- **System Discovery**: DISABLED for maximum security
- **Security Level**: **Maximum** - Zero system interaction
- **Use Cases**:
  - New users who need reliable, predictable behavior
  - Development environments where safety is paramount
  - Corporate environments with strict security requirements
  - General daily usage where built-in commands are sufficient
- **Best Practice**: This is the recommended mode for most users

### **Discovery Mode** (`what-cmd --enable-discovery`) - **EXPLICIT OPT-IN**
- **System Discovery**: ENABLED with comprehensive safety controls
- **Security Level**: **High** - Extensive safety controls with whitelist approach
- **Use Cases**:
  - Advanced users who need system integration
  - Development environments requiring custom tool discovery
  - Users who want to discover their installed CLI tools safely
- **Safety Features**:
  - Dangerous executable blacklist
  - Strict timeout controls (1-second max per help attempt)
  - Resource limits (max 50 executables per directory)
  - Process isolation with empty environment
  - Whitelist-only approach for help text extraction

### **No Discovery Mode** (`what-cmd -no-discovery`) - **EXPLICIT SAFE**
- **System Discovery**: DISABLED with explicit messaging
- **Security Level**: **Maximum** - Zero system interaction
- **Use Cases**:
  - Production environments where only known commands are needed
  - CI/CD pipelines requiring deterministic behavior
  - Restricted environments with security policies against system scanning
  - Performance-critical scenarios where discovery overhead is unacceptable
- **Identical to Safe Mode** but with explicit intent signaling

### **Emergency Mode** (`what-cmd -emergency`) - **INCIDENT RESPONSE**
- **System Discovery**: DISABLED with error signaling
- **Security Level**: **Maximum** - Complete operational shutdown
- **Use Cases**:
  - Incident response when security breach is suspected
  - Highly regulated environments (financial, healthcare, government)
  - Forensic analysis where system state must remain unchanged
  - Recovery situations where system stability is compromised
- **Behavior**: Returns explicit error messages indicating emergency shutdown

### **Command Reference**

```bash
# Safe modes (DEFAULT - no system discovery)
what-cmd                    # Safe mode (default)
what-cmd -safe             # Explicit safe mode
what-cmd -no-discovery     # Explicit no-discovery mode

# Discovery mode (EXPLICIT OPT-IN - with safety controls)
what-cmd --enable-discovery # Enable system discovery safely

# Emergency mode (INCIDENT RESPONSE)
what-cmd -emergency        # Complete shutdown with error signaling

# Other options
what-cmd -refresh          # Refresh cache (safe mode only)
```

### **Best Practices**

1. **Always start with safe mode**: `what-cmd` (default)
2. **Only enable discovery when needed**: `what-cmd --enable-discovery`
3. **Use emergency mode in sensitive environments**: `what-cmd -emergency`
4. **Test discovery in isolated environments first** before using on production systems
5. **Review logs carefully** when using discovery mode

### Interactive Features

When running what-cmd:

1. **Search**: Start typing to find commands, flags, or hotkeys
2. **Navigate**: Use ↑/↓ arrows to browse results
3. **Select**: Press Enter to select and display detailed information
4. **Exit**: Press Escape or Ctrl+C to quit

The interface shows:
- **Commands Panel**: List of matching commands
- **Descriptions Panel**: Detailed descriptions of selected items
- **Flags Panel**: Relevant flags for the selected command
- **Search Input**: Your current search query

### System Discovery

When enabled, what-cmd can discover:

- **System Commands**: CLI tools installed in your PATH
- **Custom Tools**: Applications in user-specified directories
- **Shell Aliases**: Your custom command aliases from `.bashrc`, `.zshrc`, etc.
- **Shell Functions**: Custom functions defined in shell configuration files
- **Package Manager Tools**: Commands installed via npm, pip, brew, etc.

### Configuration

Create a configuration file at `~/.what-cmd/config.json`:

```json
{
  "discovery": {
    "enabled": false,
    "scan_path": true,
    "scan_common_paths": true,
    "scan_shell_configs": true,
    "custom_paths": [
      "/opt/my-tools/bin",
      "~/personal-scripts"
    ],
    "custom_config_files": [
      "~/.my_aliases",
      "~/.custom_functions"
    ],
    "exclude_patterns": [
      "test*",
      "*.tmp",
      "*debug*"
    ],
    "max_executables": 1000,
    "refresh_interval_hours": 24
  },
  "ui": {
    "show_system_commands": true,
    "show_builtin_commands": true,
    "show_user_aliases": true,
    "group_by_source": false
  },
  "cache": {
    "enabled": true,
    "directory": "~/.what-cmd/cache",
    "max_size_mb": 50,
    "ttl_hours": 168
  }
}
```

## What's Included

### Built-in Knowledge Base

**Commands**: 500+ common terminal commands including:
- File operations (`ls`, `cp`, `mv`, `rm`)
- Text processing (`grep`, `sed`, `awk`, `sort`)
- Network tools (`curl`, `wget`, `ssh`, `scp`)
- Development tools (`git`, `npm`, `pip`, `docker`)
- System administration (`ps`, `top`, `chmod`, `chown`)

**Flags**: Comprehensive flag documentation for major commands:
- Git operations and options
- Docker container management
- Package manager flags
- File system utilities
- Network configuration

**Hotkeys**: Keyboard shortcuts for:
- **Terminal**: Bash, Zsh, PowerShell shortcuts
- **Editors**: Vim, Nano, VSCode keybindings
- **Terminal Multiplexers**: Tmux, Screen commands
- **Window Managers**: i3, Awesome, system shortcuts
- **Development**: IDE and editor shortcuts

### System Integration

**Discovered Content** (when enabled):
- Your installed CLI tools and their documentation
- Personal aliases and custom commands
- Shell functions and scripts
- Package manager installed tools
- Custom application shortcuts

### Architecture

```
what-cmd/
├── main.go                          # Application entry point
├── commands/                        # Built-in command definitions
├── flags/                          # Built-in flag definitions  
├── hotkeys/                        # Built-in hotkey definitions
└── internal/
    ├── config/                     # Configuration management
    ├── discovery/                  # System discovery engine
    │   ├── scanner.go             # Core discovery logic
    │   ├── cache.go               # Caching system
    │   └── emergency_scanner.go   # Safety wrapper
    ├── models/                     # Data structures
    ├── search/                     # Search and ranking algorithms
    └── ui/                         # Terminal user interface
```

## Troubleshooting

### Common Issues

**Permission Errors:**
```bash
# Linux/macOS: Ensure proper permissions
chmod +x what-cmd
sudo mv what-cmd /usr/local/bin/
```

**Windows PATH Issues:**
- Verify `C:\Tools` is in your PATH environment variable
- Restart terminal after PATH changes
- Use full path if needed: `C:\Tools\what-cmd.exe`

**System Discovery Problems:**
```bash
# Run in safe mode
what-cmd -safe

# Clear cache and retry
what-cmd -refresh -safe

# Use emergency mode if issues persist
what-cmd -emergency
```

**Build Issues:**
```bash
# Update Go modules
go mod tidy
go mod download

# Clean build
go clean -cache
go build -o what-cmd main.go
```

## Support

- **Issues**: Report bugs on [GitHub Issues](https://github.com/OpusMag/what-cmd/issues)

## License

This project is licensed under the MIT License