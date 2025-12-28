# what-cmd

What-cmd is a cross-platform (Windows, Linux, and macOS) command-line tool written in Go that provides an intuitive way to discover and search for terminal commands, flags, and hotkeys. Whether you've forgotten a command or you're new to the terminal, what-cmd makes it easy to find what you need. The tool is inspired by which-key and lazygit, featuring both built-in knowledge and dynamic system discovery.

![what-cmd-demo](https://github.com/user-attachments/assets/4261dfce-455f-4f73-8115-57b4c07a32d2)

## Features

### **Multi-Mode Search**
- **Commands**: Search through common terminal commands
- **Flags**: Find command-line flags and options
- **Hotkeys**: Discover keyboard shortcuts for various applications

### **System Discovery**
- **Multi-Source Architecture**: Parallel scanning of installation paths, shell configs, and user directories
- **Intelligent Caching**: Configurable TTL with automatic deduplication and performance optimization
- **Shell Integration**: Discovers aliases and functions from both user and system-wide shell configurations
- **Safe by Default**: Comprehensive security controls with multiple validation layers
- **Windows Safe Mode**: Automatic whitelist-only scanning with enhanced security on Windows
- **Legacy Support**: Backward compatibility with older configuration formats
- **Parallel Processing**: Concurrent discovery with configurable resource limits
- **Graceful Degradation**: Continues operation even if some discovery sources fail

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
# Default mode - discovers from configured paths + built-in commands
what-cmd

# Safe mode - built-in commands only (no system discovery)
what-cmd --no-discovery

# Search for flags
what-cmd --flags

# Search for hotkeys
what-cmd --hotkeys

# Force refresh discovery cache
what-cmd --refresh

# Use custom configuration file
what-cmd --config=/path/to/config.json

# Help
what-cmd --help
```

### **Command Reference**

```bash
# Default behavior
what-cmd                    # Uses discovery if installation/user paths are configured; otherwise built-in only

# Discovery controls
what-cmd --no-discovery     # Disable discovery; built-in only
what-cmd --refresh          # Refresh discovery cache (no effect if discovery is disabled)

# Configuration
what-cmd --config=./config.json
what-cmd --help
```

Note: --enable-discovery is deprecated; discovery is automatically enabled when installation_paths or user_paths are configured.

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

### Discovery

The new implementation only scans predefined paths from the user config

1. **Configuration-Driven Security**: If you configure paths, you want them scanned
2. **Safe Defaults**: No configuration = no discovery, maximum safety
3. **Explicit Override**: `--no-discovery` if you don't want discovery for some reason

### Configuration

Create a configuration file at `~/.what-cmd/config.json`:

// Replace the configuration examples with modern format:

#### Recommended config
```json
{
  "discovery": {
    "enabled": false,
    "scan_shell_configs": true,
    "installation_paths": [
      "/usr/local/bin",
      "/opt/homebrew/bin",
      "~/.local/bin",
      "~/bin",
      "~/sbin"
    ],
    "user_paths": [
      "~/personal-scripts",
      "~/dev-tools/bin"
    ],
    "exclude_patterns": [
      "*.so", "*.so.*", "*.a", "*.o",
      "test*", "*.tmp", "*debug*"
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
    "ttl_hours": 168
  }
}
```

#### Legacy Configuration (Backward Compatible)
```json
{
  "discovery": {
    "enabled": false,
    // Legacy options (still supported but deprecated)
    "scan_path": true,
    "scan_common_paths": true, 
    "custom_paths": ["/opt/my-tools"],
    "custom_config_files": ["~/.my_aliases"],
    
    // Modern equivalents (recommended)
    "installation_paths": ["/usr/local/bin", "/opt/homebrew/bin"],
    "user_paths": ["~/personal-scripts"]
  }
}
```

#### Windows-Specific Configuration
```json
{
  "discovery": {
    "enabled": false,
    "scan_shell_configs": true,
    "installation_paths": [
      "C:\\ProgramData\\chocolatey\\bin",
      "C:\\Program Files\\Git\\usr\\bin"
    ],
    "user_paths": [
      "C:\\Users\\%USERNAME%\\bin",
      "C:\\Tools"
    ],
    "windows_safe_mode": true,
    "windows_blocked_paths": [
      "C:\\Windows",
      "C:\\Program Files\\Windows*",
      "C:\\ProgramData\\Microsoft"
    ],
    "exclude_patterns": [
      "*.dll", "*.pdb", "*.lib",
      "test*", "*.tmp", "*debug*"
    ],
    "max_executables": 100,
    "refresh_interval_hours": 24
  }
}
```

#### Configuration Options Explained

**Discovery Options:**
- `enabled`: Enable/disable system discovery
- `scan_shell_configs`: Parse shell config files for aliases and functions
- `installation_paths`: Directories where CLI tools are installed (replaces `scan_path`)
- `user_paths`: User-specific directories to scan (replaces `custom_paths`)
- `exclude_patterns`: File patterns to exclude from scanning
- `max_executables`: Maximum number of executables to discover
- `refresh_interval_hours`: Cache validity period

**Legacy Options (Deprecated but Supported):**
- `scan_path`: Scan PATH environment variable (use `installation_paths` instead)
- `scan_common_paths`: Scan common installation directories (use `installation_paths`)
- `custom_paths`: Custom directories to scan (use `user_paths`)
- `custom_config_files`: Custom shell config files (automatically detected)

**Windows-Specific Options:**
- `windows_safe_mode`: Enable Windows-specific safety controls (auto-enabled)
- `windows_blocked_paths`: Blacklist of directories blocked on Windows
- `exclude_patterns`: File patterns to exclude (Windows-specific defaults)

**Performance Options:**
- `max_executables`: Limit total discovered executables
- `refresh_interval_hours`: How often to refresh the discovery cache

## Migration Guide

### From Legacy Configuration

If you're using an older configuration format, here's how to migrate:

#### Configuration Format Changes

**Legacy → Modern:**
```json
// OLD (still supported)
{
  "scan_path": true,
  "scan_common_paths": true,
  "custom_paths": ["/opt/tools", "~/scripts"],
  "custom_config_files": ["~/.my_aliases"]
}

// NEW (recommended)
{
  "installation_paths": ["/usr/local/bin", "/opt/homebrew/bin"],
  "user_paths": ["/opt/tools", "~/scripts"]
  // custom_config_files automatically detected
}
```

#### Benefits of Modern Configuration
- **Explicit Control**: Specify exactly which directories to scan
- **Better Performance**: Avoid scanning large PATH directories
- **Enhanced Security**: More granular control over discovery sources
- **Automatic Detection**: Shell config files found automatically
- **Cross-Platform**: Consistent behavior across Windows, Linux, and macOS

#### Migration Steps
1. **Backup** your current `~/.what-cmd/config.json`
2. **Update** configuration format using examples above
3. **Test** with `what-cmd --refresh` (discovery runs automatically if paths are configured)
4. **Verify** that all your custom tools are still discovered

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
├── main.go                         # Application entry point
├── commands/                       # Built-in command definitions
├── flags/                          # Built-in flag definitions  
├── hotkeys/                        # Built-in hotkey definitions
└── internal/
    ├── config/                     # Configuration management
    ├── discovery/                  # System discovery engine
    │   ├── scanner.go              # Core discovery logic with security
    │   ├── cache.go                # Caching system
    ├── models/                     # Data structures
    ├── search/                     # Search and ranking algorithms
    └── ui/                         # Terminal user interface
```

### Discovery Engine Architecture

```
SystemScanner
├── Multi-Source Discovery (Parallel)
│   ├── Installation Paths Scanner
│   ├── Shell Config Parser
│   ├── User Paths Scanner
│   └── Legacy PATH Scanner (deprecated)
├── Security Layer
│   ├── Path Validation
│   ├── Executable Filtering  
│   ├── Timeout Controls
│   └── Resource Limits
├── Caching System
│   ├── Intelligent TTL
│   ├── Deduplication
│   └── Performance Optimization
└── Emergency Controls
    ├── Safe Mode Enforcement
    ├── Windows-Specific Protections
    └── Incident Response Mode
```

#### **Safety Features**
- **Concurrency Control**: Semaphore-based limiting of concurrent directory scans
- **Description Extraction**: Safe help text extraction with multiple fallback strategies
- **Pattern Matching**: Intelligent command description inference for common tools
- **Error Handling**: Graceful degradation when discovery sources fail

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
what-cmd --no-discovery

# Clear cache and retry (discovery will run if paths are configured)
what-cmd --refresh

# Check configuration file
what-cmd --config=./config.json --refresh
```

**Windows-Specific Issues:**
- **"SECURITY: Directory access denied"**: This is normal - Windows safe mode blocks system directories
- **Limited results on Windows**: By design - only user-configured paths are scanned
- **Configuration needed**: Add safe directories to user_paths in the config file (custom_paths is legacy)

**Build Issues:**
```bash
# Update Go modules
go mod tidy
go mod download

# Clean build
go clean -cache
go build -o what-cmd main.go
```
**Discovery Configuration Issues:**
```bash
# Check if you're using legacy configuration
what-cmd --config=./config.json

# Migrate to modern configuration format
# Replace "custom_paths" with "user_paths"
# Replace "scan_path: true" with specific "installation_paths"

# Debug discovery sources (forces cache refresh)
what-cmd --refresh
```