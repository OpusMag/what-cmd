package discovery

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"sync"
	"time"

	"what-cmd/internal/config"
	"what-cmd/internal/models"
)

type SystemScanner struct {
	config          *config.Config
	cache           *Cache
	mu              sync.RWMutex
	lastScan        time.Time
	discoveredItems []models.Item
	safetyLimits    SafetyLimits
}

type SafetyLimits struct {
	MaxConcurrentScans   int
	MaxHelpAttempts      int
	HelpCommandTimeout   time.Duration
	MaxExecutablesPerDir int
	DangerousExecutables []string
}

func DefaultSafetyLimits() SafetyLimits {
	limits := SafetyLimits{
		MaxConcurrentScans:   4,
		MaxHelpAttempts:      25,
		HelpCommandTimeout:   500 * time.Millisecond,
		MaxExecutablesPerDir: 50,
		DangerousExecutables: []string{
			"AppHostNameRegistrationVerifier.exe",
			"SystemSettingsAdminFlows.exe",
			"SystemSettingsRemoveDevice.exe",
			"WindowsActionDialog.exe",
			"UserAccountControlSettings.exe",
			"shutdown.exe", "reboot.exe", "format.exe",
			"del.exe", "rmdir.exe", "diskpart.exe",
			"regedit.exe", "reg.exe", "powercfg.exe",
			"bcdedit.exe", "dism.exe", "sfc.exe",
			"chkdsk.exe", "msiexec.exe",
		},
	}

	if runtime.GOOS == "windows" {
		limits.MaxHelpAttempts = 10
		limits.HelpCommandTimeout = 300 * time.Millisecond
		limits.MaxExecutablesPerDir = 25
	}

	return limits
}

func NewSystemScanner(cfg *config.Config) (*SystemScanner, error) {
	cache, err := NewCache(cfg.Cache)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize cache: %w", err)
	}

	scanner := &SystemScanner{
		config:       cfg,
		cache:        cache,
		safetyLimits: DefaultSafetyLimits(),
	}

	return scanner, nil
}

func (s *SystemScanner) getDiscoveryPaths() []string {
	effectivePaths := s.config.Discovery.GetEffectivePaths()
	return s.filterAndValidatePaths(effectivePaths)
}

func (s *SystemScanner) filterAndValidatePaths(paths []string) []string {
	seen := make(map[string]bool)
	var validPaths []string

	for _, path := range paths {
		expandedPath := s.expandPath(path)

		normalizedPath := strings.ToLower(filepath.Clean(expandedPath))
		if seen[normalizedPath] {
			continue
		}
		seen[normalizedPath] = true

		if !s.isPathAllowed(expandedPath) {
			fmt.Fprintf(os.Stderr, "SECURITY: Blocked path: %s\n", expandedPath)
			continue
		}

		if !s.isDirectoryAccessible(expandedPath) {
			continue
		}

		validPaths = append(validPaths, expandedPath)
	}

	return validPaths
}

func (s *SystemScanner) expandPath(path string) string {
	if strings.HasPrefix(path, "~/") {
		if homeDir, err := os.UserHomeDir(); err == nil {
			return filepath.Join(homeDir, path[2:])
		}
	}
	return os.ExpandEnv(path)
}

func (s *SystemScanner) isDirectoryAccessible(dirPath string) bool {
	info, err := os.Stat(dirPath)
	if err != nil {
		return false
	}

	if !info.IsDir() {
		return false
	}

	_, err = os.ReadDir(dirPath)
	return err == nil
}

func (s *SystemScanner) discoverExecutables(ctx context.Context) ([]models.Item, error) {
	discoveryPaths := s.getDiscoveryPaths()

	fmt.Fprintf(os.Stderr, "Scanning %d configured directories...\n", len(discoveryPaths))

	var allItems []models.Item
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, s.safetyLimits.MaxConcurrentScans)
	itemsChan := make(chan []models.Item, len(discoveryPaths))
	errorsChan := make(chan error, len(discoveryPaths))

	for _, path := range discoveryPaths {
		wg.Add(1)
		go func(dirPath string) {
			defer wg.Done()

			select {
			case semaphore <- struct{}{}:
				defer func() { <-semaphore }()
			case <-ctx.Done():
				return
			}

			items, err := s.scanDirectory(dirPath, models.SystemCommand)
			if err != nil {
				errorsChan <- fmt.Errorf("failed to scan %s: %w", dirPath, err)
				return
			}

			itemsChan <- items
			if len(items) > 0 {
				fmt.Fprintf(os.Stderr, "%s: %d executables\n", dirPath, len(items))
			}
		}(path)
	}

	go func() {
		wg.Wait()
		close(itemsChan)
		close(errorsChan)
	}()

	for items := range itemsChan {
		allItems = append(allItems, items...)
	}

	for err := range errorsChan {
		fmt.Fprintf(os.Stderr, "⚠️  Warning: %v\n", err)
	}

	return allItems, nil
}

func (s *SystemScanner) DiscoverItems(ctx context.Context) ([]models.Item, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.shouldUseCachedResults() {
		if items, err := s.cache.GetItems(); err == nil && len(items) > 0 {
			s.discoveredItems = items
			fmt.Fprintf(os.Stderr, "📋 Using cached discovery results: %d items\n", len(items))
			return items, nil
		}
	}

	fmt.Fprintf(os.Stderr, "🔍 Starting enhanced system discovery...\n")
	if runtime.GOOS == "windows" && s.config.Discovery.WindowsSafeMode {
		fmt.Fprintf(os.Stderr, "🛡️  Windows Safe Mode: Using predefined installation paths only\n")
	}

	var allItems []models.Item
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, s.safetyLimits.MaxConcurrentScans)
	itemsChan := make(chan []models.Item, 4)
	errorsChan := make(chan error, 4)

	discoveryCtx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	discoveryTasks := []struct {
		name     string
		enabled  bool
		taskFunc func(context.Context) ([]models.Item, error)
	}{
		{"PredefinedPaths", true, s.discoverExecutables},
		{"ShellConfigs", s.config.Discovery.ScanShellConfigs, s.scanShellConfigs},
		{"LegacyPATH", s.config.Discovery.ScanPATH, s.scanPATHExecutables},      // Legacy support
		{"LegacyCommon", s.config.Discovery.ScanCommonPaths, s.scanCommonPaths}, // Legacy support
	}

	for _, task := range discoveryTasks {
		if task.enabled {
			wg.Add(1)
			go s.safeExecuteDiscovery(&wg, semaphore, discoveryCtx, itemsChan, errorsChan, task.taskFunc)
		}
	}

	go func() {
		wg.Wait()
		close(itemsChan)
		close(errorsChan)
	}()

	for items := range itemsChan {
		allItems = append(allItems, items...)

		if len(allItems) > s.config.Discovery.MaxExecutables*3 {
			fmt.Fprintf(os.Stderr, "Discovery safety limit reached, truncating results\n")
			break
		}
	}

	for err := range errorsChan {
		fmt.Fprintf(os.Stderr, "Discovery warning: %v\n", err)
	}

	allItems = s.deduplicateItems(allItems)
	if len(allItems) > s.config.Discovery.MaxExecutables {
		fmt.Fprintf(os.Stderr, "Limiting results to %d executables (found %d)\n",
			s.config.Discovery.MaxExecutables, len(allItems))
		allItems = allItems[:s.config.Discovery.MaxExecutables]
	}

	if s.config.Cache.Enabled {
		if err := s.cache.StoreItems(allItems); err != nil {
			fmt.Fprintf(os.Stderr, "Cache storage warning: %v\n", err)
		}
	}

	s.discoveredItems = allItems
	s.lastScan = time.Now()

	fmt.Fprintf(os.Stderr, "Enhanced discovery completed: %d items found\n", len(allItems))
	return allItems, nil
}

// Legacy support methods (maintained for backward compatibility)
func (s *SystemScanner) scanPATHExecutables(ctx context.Context) ([]models.Item, error) {
	if runtime.GOOS == "windows" && s.config.Discovery.WindowsSafeMode {
		fmt.Fprintf(os.Stderr, "Windows Safe Mode: Skipping legacy PATH scan for security\n")
		return []models.Item{}, nil
	}

	fmt.Fprintf(os.Stderr, "DEPRECATED: Using legacy PATH scanning. Consider migrating to 'installation_paths' configuration.\n")

	pathEnv := os.Getenv("PATH")
	if pathEnv == "" {
		return nil, nil
	}

	var items []models.Item
	pathSeparator := ":"
	if runtime.GOOS == "windows" {
		pathSeparator = ";"
	}

	paths := strings.Split(pathEnv, pathSeparator)

	for _, path := range paths {
		select {
		case <-ctx.Done():
			return items, ctx.Err()
		default:
		}

		if path == "" {
			continue
		}

		if !s.isPathAllowed(path) {
			fmt.Fprintf(os.Stderr, "SECURITY: Skipping disallowed PATH directory: %s\n", path)
			continue
		}

		pathItems, err := s.scanDirectory(path, models.SystemCommand)
		if err != nil {
			continue
		}

		items = append(items, pathItems...)
	}

	return items, nil
}

func (s *SystemScanner) scanCommonPaths(ctx context.Context) ([]models.Item, error) {
	if runtime.GOOS == "windows" && s.config.Discovery.WindowsSafeMode {
		fmt.Fprintf(os.Stderr, "Windows Safe Mode: Skipping legacy common paths scan for security\n")
		return []models.Item{}, nil
	}

	fmt.Fprintf(os.Stderr, "DEPRECATED: Using legacy common paths scanning. Consider migrating to 'installation_paths' configuration.\n")

	var items []models.Item
	commonPaths := s.getCommonInstallationPaths()

	for _, path := range commonPaths {
		select {
		case <-ctx.Done():
			return items, ctx.Err()
		default:
		}

		if !s.isPathAllowed(path) {
			fmt.Fprintf(os.Stderr, "SECURITY: Skipping disallowed common directory: %s\n", path)
			continue
		}

		pathItems, err := s.scanDirectory(path, models.SystemCommand)
		if err != nil {
			continue
		}

		items = append(items, pathItems...)
	}

	return items, nil
}

func (s *SystemScanner) shouldUseCachedResults() bool {
	if !s.config.Cache.Enabled {
		return false
	}

	refreshInterval := time.Duration(s.config.Discovery.RefreshIntervalHours) * time.Hour
	return time.Since(s.lastScan) < refreshInterval
}

func (s *SystemScanner) shouldExcludeFile(filename string) bool {
	for _, pattern := range s.config.Discovery.ExcludePatterns {
		if matched, _ := filepath.Match(pattern, filename); matched {
			return true
		}
	}
	return false
}

func (s *SystemScanner) isExecutable(filePath string) bool {
	info, err := os.Stat(filePath)
	if err != nil {
		return false
	}

	mode := info.Mode()

	if runtime.GOOS == "windows" {
		ext := strings.ToLower(filepath.Ext(filePath))
		executableExts := []string{".exe", ".bat", ".cmd", ".com", ".ps1"}
		for _, execExt := range executableExts {
			if ext == execExt {
				return true
			}
		}
		return false
	}

	return mode&0111 != 0
}

func (s *SystemScanner) getDefaultConfigFiles() []string {
	homeDir, _ := os.UserHomeDir()

	switch runtime.GOOS {
	case "windows":
		return []string{
			filepath.Join(homeDir, "Documents", "PowerShell", "profile.ps1"),
			filepath.Join(homeDir, ".gitconfig"),
		}
	default:
		return []string{
			filepath.Join(homeDir, ".bashrc"),
			filepath.Join(homeDir, ".bash_profile"),
			filepath.Join(homeDir, ".zshrc"),
			filepath.Join(homeDir, ".profile"),
			filepath.Join(homeDir, ".bash_aliases"),
			filepath.Join(homeDir, ".gitconfig"),
			filepath.Join(homeDir, ".config", "fish", "config.fish"),
			"/etc/bash.bashrc",
			"/etc/bashrc",
			"/etc/profile",
			"/etc/zsh/zshrc",
		}
	}
}

func (s *SystemScanner) parseShellConfig(configPath string) ([]models.Item, error) {
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var items []models.Item
	scanner := bufio.NewScanner(file)

	aliasPattern := regexp.MustCompile(`^\s*alias\s+([^=\s]+)=(.*)$`)
	functionPattern := regexp.MustCompile(`^\s*function\s+([^\s\(]+)`)
	bashFunctionPattern := regexp.MustCompile(`^\s*([a-zA-Z_][a-zA-Z0-9_]*)\s*\(\s*\)\s*\{`)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		if matches := aliasPattern.FindStringSubmatch(line); matches != nil {
			aliasName := matches[1]
			aliasValue := strings.Trim(matches[2], `"'`)

			items = append(items, models.Item{
				Name:        aliasName,
				Description: fmt.Sprintf("Alias for: %s", aliasValue),
				Type:        models.UserAlias,
				Source:      configPath,
			})
		}

		if matches := functionPattern.FindStringSubmatch(line); matches != nil {
			functionName := matches[1]
			items = append(items, models.Item{
				Name:        functionName,
				Description: "User-defined shell function",
				Type:        models.UserFunction,
				Source:      configPath,
			})
		}

		if matches := bashFunctionPattern.FindStringSubmatch(line); matches != nil {
			functionName := matches[1]
			items = append(items, models.Item{
				Name:        functionName,
				Description: fmt.Sprintf("Shell function: %s", functionName),
				Type:        models.UserFunction,
				Source:      configPath,
			})
		}
	}

	return items, scanner.Err()
}

func (s *SystemScanner) getManPageDescription(cmdName string) string {
	if runtime.GOOS == "windows" {
		return ""
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "man", "-f", cmdName)
	if output, err := cmd.Output(); err == nil {
		lines := strings.Split(string(output), "\n")
		for _, line := range lines {
			if strings.Contains(line, " - ") {
				parts := strings.SplitN(line, " - ", 2)
				if len(parts) == 2 {
					return strings.TrimSpace(parts[1])
				}
			}
		}
	}

	return ""
}

func (s *SystemScanner) inferDescriptionFromName(cmdName string) string {
	namePatterns := map[string]string{
		"git":     "Version control system",
		"docker":  "Container management",
		"kubectl": "Kubernetes control",
		"npm":     "Node.js package manager",
		"pip":     "Python package installer",
		"cargo":   "Rust package manager",
		"go":      "Go programming language tool",
		"python":  "Python interpreter",
		"node":    "Node.js runtime",
		"java":    "Java runtime",
		"gcc":     "GNU Compiler Collection",
		"clang":   "C language family frontend",
		"make":    "Build automation tool",
		"cmake":   "Cross-platform build system",
		"tldr":    "Simplified command examples",
		"bat":     "Enhanced cat with syntax highlighting",
		"fd":      "Fast alternative to find",
		"rg":      "Fast text search tool",
		"exa":     "Modern ls replacement",
	}

	lowerName := strings.ToLower(cmdName)
	for pattern, desc := range namePatterns {
		if strings.Contains(lowerName, pattern) {
			return desc
		}
	}

	return "System command"
}

func (s *SystemScanner) deduplicateItems(items []models.Item) []models.Item {
	seen := make(map[string]bool)
	var unique []models.Item

	for _, item := range items {
		key := fmt.Sprintf("%s:%d", item.Name, item.Type)
		if !seen[key] {
			seen[key] = true
			unique = append(unique, item)
		}
	}

	return unique
}

func (s *SystemScanner) isPathAllowed(dirPath string) bool {
	if runtime.GOOS != "windows" {
		return true
	}

	if !s.config.Discovery.WindowsSafeMode {
		return true
	}

	lowerPath := strings.ToLower(filepath.Clean(dirPath))

	for _, blockedPath := range s.config.Discovery.WindowsBlockedPaths {
		blockedLower := strings.ToLower(filepath.Clean(blockedPath))
		if strings.HasPrefix(lowerPath, blockedLower) {
			return false
		}
	}

	effectivePaths := s.config.Discovery.GetEffectivePaths()
	for _, allowedPath := range effectivePaths {
		allowedLower := strings.ToLower(filepath.Clean(s.expandPath(allowedPath)))
		if strings.HasPrefix(lowerPath, allowedLower) {
			return true
		}
	}

	return false
}

func (s *SystemScanner) scanShellConfigs(ctx context.Context) ([]models.Item, error) {
	var items []models.Item

	configFiles := append(s.config.Discovery.CustomConfigFiles,
		s.getDefaultConfigFiles()...)

	for _, configFile := range configFiles {
		select {
		case <-ctx.Done():
			return items, ctx.Err()
		default:
		}

		configItems, err := s.parseShellConfig(configFile)
		if err != nil {
			continue
		}

		items = append(items, configItems...)
	}

	return items, nil
}

func (s *SystemScanner) scanDirectory(dirPath string, itemType models.ItemType) ([]models.Item, error) {
	if !s.isPathAllowed(dirPath) {
		fmt.Fprintf(os.Stderr, "SECURITY: Directory access denied: %s\n", dirPath)
		return []models.Item{}, nil
	}

	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		return []models.Item{}, nil
	} else if err != nil {
		return nil, fmt.Errorf("cannot access directory %s: %w", dirPath, err)
	}

	var items []models.Item
	helpAttemptCount := 0

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory %s: %w", dirPath, err)
	}

	processedCount := 0
	for _, entry := range entries {
		if processedCount >= s.safetyLimits.MaxExecutablesPerDir {
			fmt.Fprintf(os.Stderr, "SAFETY: Reached max executables limit for directory: %s\n", dirPath)
			break
		}

		if entry.IsDir() {
			continue
		}

		filename := entry.Name()

		if s.isDangerousExecutable(filename) {
			fmt.Fprintf(os.Stderr, "SAFETY: Blocked dangerous executable: %s in %s\n", filename, dirPath)
			continue
		}

		fullPath := filepath.Join(dirPath, filename)

		if s.shouldExcludeFile(filename) {
			continue
		}

		if !s.isExecutable(fullPath) {
			continue
		}

		var description string
		if s.shouldSkipHelpAttempt(filename) {
			description = s.inferDescriptionFromName(filename)
		} else if helpAttemptCount < s.safetyLimits.MaxHelpAttempts {
			description = s.getSafeCommandDescription(filename, fullPath)
			if description != "" {
				helpAttemptCount++
			}
		} else {
			description = s.inferDescriptionFromName(filename)
		}

		items = append(items, models.Item{
			Name:        filename,
			Description: description,
			Type:        itemType,
			Source:      dirPath,
		})

		processedCount++
	}

	return items, nil
}

func (s *SystemScanner) safeExecuteDiscovery(
	wg *sync.WaitGroup,
	semaphore chan struct{},
	ctx context.Context,
	itemsChan chan<- []models.Item,
	errorsChan chan<- error,
	discoveryFunc func(context.Context) ([]models.Item, error),
) {
	defer wg.Done()

	select {
	case semaphore <- struct{}{}:
		defer func() { <-semaphore }()
	case <-ctx.Done():
		return
	}

	func() {
		defer func() {
			if r := recover(); r != nil {
				errorsChan <- fmt.Errorf("discovery panic recovered: %v", r)
			}
		}()

		items, err := discoveryFunc(ctx)
		if err != nil {
			errorsChan <- err
		} else {
			itemsChan <- items
		}
	}()
}

func (s *SystemScanner) isDangerousExecutable(filename string) bool {
	lowerName := strings.ToLower(filename)

	for _, dangerous := range s.safetyLimits.DangerousExecutables {
		if lowerName == strings.ToLower(dangerous) {
			return true
		}
	}

	dangerousPatterns := []string{
		"systemsettings", "useraccount", "apphostname", "windowsaction",
		"adminflows", "removedevice", "registrationverifier",
		"shutdown", "format", "diskpart", "regedit", "powercfg",
		"bcdedit", "dism", "sfc", "chkdsk", "microsoft", "windows",
		"setup", "install", "uninstall", "msiexec", "config",
		"admin", "control", "manage",
	}

	for _, pattern := range dangerousPatterns {
		if strings.Contains(lowerName, pattern) {
			return true
		}
	}

	return false
}

func (s *SystemScanner) shouldSkipHelpAttempt(filename string) bool {
	lower := strings.ToLower(filename)

	neverExecute := []string{
		"apphostname", "systemsettings", "useraccount", "windowsaction",
		"adminflows", "removedevice", "registrationverifier",
		"shutdown", "format", "diskpart", "regedit", "powercfg",
		"bcdedit", "dism", "sfc", "chkdsk", "msiexec",
		"setup", "install", "uninstall", "config", "admin",
		"control", "manage", "microsoft", "windows",
		".dll", ".sys", ".msi", ".bat", ".cmd", ".ps1", ".vbs", ".js",
	}

	for _, pattern := range neverExecute {
		if strings.Contains(lower, pattern) {
			return true
		}
	}

	return false
}

func (s *SystemScanner) getSafeCommandDescription(cmdName, cmdPath string) string {
	if s.shouldSkipHelpAttempt(cmdName) {
		fmt.Fprintf(os.Stderr, "SAFETY: Blocked help attempt for: %s\n", cmdName)
		return s.inferDescriptionFromName(cmdName)
	}

	if s.isDangerousExecutable(filepath.Base(cmdPath)) {
		fmt.Fprintf(os.Stderr, "SAFETY: Blocked dangerous path: %s\n", cmdPath)
		return s.inferDescriptionFromName(cmdName)
	}

	if desc := s.getSafeHelpFromCommand(cmdPath); desc != "" {
		return desc
	}

	if desc := s.getManPageDescription(cmdName); desc != "" {
		return desc
	}

	return s.inferDescriptionFromName(cmdName)
}

func (s *SystemScanner) getSafeHelpFromCommand(cmdPath string) string {
	filename := filepath.Base(cmdPath)
	if s.isDangerousExecutable(filename) || s.shouldSkipHelpAttempt(filename) {
		fmt.Fprintf(os.Stderr, "SAFETY: Final block for: %s\n", filename)
		return ""
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.safetyLimits.HelpCommandTimeout)
	defer cancel()

	safeHelpFlags := []string{"--help", "-h"}

	for _, flag := range safeHelpFlags {
		select {
		case <-ctx.Done():
			return ""
		default:
		}

		cmd := exec.CommandContext(ctx, cmdPath, flag)
		cmd.Env = []string{"PATH="}

		output, err := cmd.CombinedOutput()
		if err != nil {
			continue
		}

		outputStr := string(output)
		if len(outputStr) > 10000 {
			outputStr = outputStr[:10000]
		}

		lines := strings.Split(outputStr, "\n")
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if len(line) > 10 && len(line) < 200 &&
				!strings.Contains(strings.ToLower(line), "usage:") &&
				!strings.Contains(strings.ToLower(line), "error:") &&
				!strings.Contains(strings.ToLower(line), "fatal:") {
				return line
			}
		}
	}

	return ""
}

// Legacy method (maintained for backward compatibility)
func (s *SystemScanner) getCommonInstallationPaths() []string {
	switch runtime.GOOS {
	case "windows":
		return []string{
			`C:\ProgramData\chocolatey\bin`,
		}
	case "darwin":
		return []string{
			"/usr/local/bin",
			"/opt/homebrew/bin",
		}
	default:
		return []string{
			"/usr/bin",
			"/usr/local/bin",
			"/snap/bin",
		}
	}
}

func (s *SystemScanner) ClearCache() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.discoveredItems = nil
	s.lastScan = time.Time{}

	return s.cache.ClearCache()
}

func (s *SystemScanner) EmergencyStop() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.discoveredItems = nil
	s.lastScan = time.Time{}

	fmt.Fprintf(os.Stderr, "Emergency stop activated - discovery halted\n")
}
