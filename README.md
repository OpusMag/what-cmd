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
- **Windows Safe Mode**: Special security controls for Windows with whitelist-only directory scanning
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
what-cmd --flags

# Search for hotkeys
what-cmd --hotkeys

# Use your custom configuration file
what-cmd --config=/path/to/config.json
```

## Security & Safety

Security is a high priority in what-cmd's design. When implementing system discovery, I have tried to make the implementation as safe as possible. This includes preventing it from accessing or manipulating files and directories it shouldn't. This is especially the case on windows, because during development, there were bugs where it did those things. If you encounter anything like that, on any platform, please raise an issue as this is a problem I do not want in what-cmd.

**What-cmd offers four distinct security modes:**

### **Safe Mode** (`what-cmd` or `what-cmd --safe`) - **DEFAULT**
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
- **Security Level**: **High** - Extensive safety controls with multiple protection layers
- **Use Cases**:
  - Advanced users who need system integration
  - Development environments requiring custom tool discovery
  - Users who want to discover their installed CLI tools safely
- **Safety Features**:
  - **Windows Safe Mode**: Automatic activation on Windows with whitelist-only scanning
  - **Dangerous Executable Blacklist**: Blocks known problematic system executables
  - **Strict Timeout Controls**: 200ms max per help attempt on Windows, 500ms on other platforms
  - **Resource Limits**: Max 10 executables per directory on Windows, 25 on other platforms
  - **Process Isolation**: Empty environment variables during help text extraction
  - **Path Validation**: Multiple checkpoints before any directory access
  - **Pattern-Based Blocking**: Prevents execution of installers, system tools, and configuration utilities

### **No Discovery Mode** (`what-cmd --no-discovery`) - **EXPLICIT SAFE**
- **System Discovery**: DISABLED with explicit messaging
- **Security Level**: **Maximum** - Zero system interaction
- **Use Cases**:
  - Production environments where only known commands are needed
  - CI/CD pipelines requiring deterministic behavior
  - Restricted environments with security policies against system scanning
  - Performance-critical scenarios where discovery overhead is unacceptable
- **Identical to Safe Mode** but with explicit intent signaling

### **Emergency Mode** (`what-cmd --emergency`) - **INCIDENT RESPONSE**
- **System Discovery**: DISABLED with error signaling
- **Security Level**: **Maximum** - Complete operational shutdown
- **Use Cases**:
  - Incident response when security breach is suspected
  - Highly regulated environments (financial, healthcare, government)
  - Forensic analysis where system state must remain unchanged
  - Recovery situations where system stability is compromised
- **Behavior**: Returns explicit error messages indicating emergency shutdown

### **Windows-Specific Security**

Windows systems have additional security measures:

- **Windows Safe Mode**: Automatically enabled during discovery on Windows
- **Blocked Directories**: System directories are completely excluded from scanning:
  - `C:\Windows\*` (all Windows system directories)
  - `C:\Program Files\Windows*`
  - `C:\ProgramData\Microsoft`
  - All Windows system subdirectories
- **Allowed Paths**: Only user-specified or safe directories are scanned:
  - `C:\ProgramData\chocolatey\bin` (default safe path)
  - User-configured paths in `custom_paths`
- **Enhanced Executable Detection**: Comprehensive blocking of:
  - System configuration tools (`regedit`, `powercfg`, `bcdedit`)
  - File system utilities (`format`, `diskpart`, `sfc`)
  - Windows-specific dangerous executables (`AppHostNameRegistrationVerifier`, etc.)
  - Installers and setup programs
- **Reduced Limits**: More conservative resource limits on Windows

### **Command Reference**

```bash
# Safe modes (DEFAULT - no system discovery)
what-cmd                    # Safe mode (default)
what-cmd --safe             # Explicit safe mode
what-cmd --no-discovery     # Explicit no-discovery mode

# Discovery mode (EXPLICIT OPT-IN - with safety controls)
what-cmd --enable-discovery # Enable system discovery safely

# Emergency mode (INCIDENT RESPONSE)
what-cmd --emergency        # Complete shutdown with error signaling

# Other options
what-cmd --refresh          # Refresh cache (works with discovery modes)
what-cmd --unsafe           # Minimal safety mode (NOT RECOMMENDED)
```

### **Best Practices**

1. **Always start with safe mode**: `what-cmd` (default)
2. **Only enable discovery when needed**: `what-cmd --enable-discovery`
3. **Configure allowed paths on Windows**: Use config file to specify safe directories
4. **Test discovery in isolated environments first** before using on production systems
5. **Review logs carefully** when using discovery mode - blocked items are logged for transparency
6. **Use emergency mode in sensitive environments**: `what-cmd --emergency`

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

- **System Commands**: CLI tools installed in your PATH (Unix/Linux/macOS)
- **Custom Tools**: Applications in user-specified directories
- **Shell Aliases**: Your custom command aliases from `.bashrc`, `.zshrc`, etc.
- **Shell Functions**: Custom functions defined in shell configuration files
- **Package Manager Tools**: Commands installed via chocolatey (Windows), homebrew (macOS), or system packages (Linux)

**Note**: On Windows, PATH scanning is disabled by default for security. Only user-configured directories are scanned.

### Configuration

Create a configuration file at `~/.what-cmd/config.json`:

#### Basic Configuration
```json
{
  "discovery": {
    "enabled": false,
    "scan_path": false,
    "scan_common_paths": false,
    "scan_shell_configs": true,
    "custom_paths": [],
    "exclude_patterns": [
      "test*",
      "*.tmp",
      "*debug*"
    ],
    "max_executables": 100,
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

#### Windows-Specific Configuration
```json
{
  "discovery": {
    "enabled": false,
    "scan_path": false,
    "scan_common_paths": false,
    "scan_shell_configs": true,
    "custom_paths": [
      "C:\\ProgramData\\chocolatey\\bin",
      "C:\\Program Files\\Git\\usr\\bin",
      "C:\\Program Files\\PowerShell\\7"
    ],
    "windows_safe_mode": true,
    "windows_allowed_paths": [
      "C:\\ProgramData\\chocolatey\\bin",
      "C:\\Program Files\\Git\\usr\\bin",
      "C:\\Program Files\\PowerShell\\7",
      "C:\\Users\\%USERNAME%\\bin"
    ],
    "windows_blocked_paths": [
      "C:\\Windows",
      "C:\\Program Files\\Windows",
      "C:\\ProgramData\\Microsoft"
    ],
    "exclude_patterns": [
      "*.dll",
      "*.pdb",
      "*.lib",
      "test*",
      "*.tmp",
      "*debug*"
    ],
    "max_executables": 100,
    "refresh_interval_hours": 24
  }
}
```

#### Linux/macOS Configuration
```json
{
  "discovery": {
    "enabled": false,
    "scan_path": true,
    "scan_common_paths": true,
    "scan_shell_configs": true,
    "custom_paths": [
      "/opt/my-tools/bin",
      "~/personal-scripts",
      "~/.local/bin"
    ],
    "exclude_patterns": [
      "*.so",
      "*.so.*",
      "*.a",
      "*.o",
      "test*",
      "*.tmp",
      "*debug*"
    ],
    "max_executables": 1000,
    "refresh_interval_hours": 24
  }
}
```

#### Configuration Options Explained

**Discovery Options:**
- `enabled`: Enable/disable system discovery
- `scan_path`: Scan directories in PATH environment variable (disabled on Windows by default)
- `scan_common_paths`: Scan common installation directories (disabled on Windows by default)
- `scan_shell_configs`: Parse shell configuration files for aliases and functions
- `custom_paths`: User-specified directories to scan
- `windows_safe_mode`: Enable Windows-specific safety controls (auto-enabled on Windows)
- `windows_allowed_paths`: Whitelist of directories allowed on Windows
- `windows_blocked_paths`: Blacklist of directories blocked on Windows
- `exclude_patterns`: File patterns to exclude from scanning
- `max_executables`: Maximum number of executables to discover
- `refresh_interval_hours`: Cache validity period

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
- Package manager installed tools (chocolatey, homebrew, apt, etc.)
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
    │   ├── scanner.go             # Core discovery logic with security
    │   ├── cache.go               # Caching system
    │   └── emergency_scanner.go   # Emergency mode wrapper
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
# Run in safe mode first
what-cmd --safe

# Clear cache and retry
what-cmd --refresh --enable-discovery

# Check configuration file
what-cmd --config=./config.json --enable-discovery

# Use emergency mode if issues persist
what-cmd --emergency
```

**Windows-Specific Issues:**
- **"SECURITY: Directory access denied"**: This is normal - Windows safe mode blocks system directories
- **Limited results on Windows**: By design - only user-configured paths are scanned
- **Configuration needed**: Add safe directories to `custom_paths` in config file

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

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.