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
	config       *config.Config
	cache        *Cache
	mu           sync.RWMutex
	safetyLimits SafetyLimits
}

type SafetyLimits struct {
	MaxConcurrentScans   int
	MaxHelpAttempts      int
	HelpCommandTimeout   time.Duration
	MaxExecutablesPerDir int
	DangerousExecutables []string
}

type DiscoveryTask struct {
	Name    string
	Enabled bool
	Execute func(context.Context) ([]models.Item, error)
}

func defaultSafetyLimits() SafetyLimits {
	limits := SafetyLimits{
		MaxConcurrentScans:   4,
		MaxHelpAttempts:      30,
		HelpCommandTimeout:   500 * time.Millisecond,
		MaxExecutablesPerDir: 10000,
		DangerousExecutables: []string{
			"shutdown.exe", "reboot.exe", "format.exe", "diskpart.exe",
			"regedit.exe", "reg.exe", "powercfg.exe", "bcdedit.exe",
			"dism.exe", "sfc.exe", "chkdsk.exe", "msiexec.exe",
			"AppHostNameRegistrationVerifier.exe",
			"SystemSettingsAdminFlows.exe",
			"SystemSettingsRemoveDevice.exe",
			"WindowsActionDialog.exe",
			"UserAccountControlSettings.exe",
		},
	}

	if runtime.GOOS == "windows" {
		limits.MaxHelpAttempts = 15
		limits.HelpCommandTimeout = 300 * time.Millisecond
		limits.MaxExecutablesPerDir = 50
	}

	return limits
}

func NewSystemScanner(cfg *config.Config) (*SystemScanner, error) {
	if cfg == nil {
		return nil, fmt.Errorf("config cannot be nil")
	}

	cache, err := NewCache(cfg.Cache)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize cache: %w", err)
	}

	return &SystemScanner{
		config:       cfg,
		cache:        cache,
		safetyLimits: defaultSafetyLimits(),
	}, nil
}

func (s *SystemScanner) DiscoverItems(ctx context.Context) ([]models.Item, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.config.Cache.Enabled {
		if cachedItems, err := s.cache.GetItems(); err == nil && len(cachedItems) > 0 {
			fmt.Fprintf(os.Stderr, "Using cached results: %d items (cache hit)\n", len(cachedItems))
			return cachedItems, nil
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "Cache miss: %v - performing fresh discovery\n", err)
		}
	}

	fmt.Fprintf(os.Stderr, "Starting path-based discovery...\n")

	discoveryCtx, cancel := context.WithTimeout(ctx, 25*time.Second)
	defer cancel()

	tasks := []DiscoveryTask{
		{
			Name:    "ConfiguredPaths",
			Enabled: len(s.config.Discovery.GetEffectivePaths()) > 0,
			Execute: s.scanConfiguredPaths,
		},
		{
			Name:    "ShellConfigs",
			Enabled: s.config.Discovery.ScanShellConfigs,
			Execute: s.scanShellConfigs,
		},
		{
			Name:    "LegacyPATH",
			Enabled: s.config.Discovery.ScanPATH,
			Execute: s.scanLegacyPATH,
		},
		{
			Name:    "LegacyCommon",
			Enabled: s.config.Discovery.ScanCommonPaths,
			Execute: s.scanLegacyCommonPaths,
		},
	}

	allItems, err := s.executeDiscoveryTasks(discoveryCtx, tasks)
	if err != nil {
		return nil, fmt.Errorf("discovery execution failed: %w", err)
	}

	allItems = s.postProcessItems(allItems)

	if s.config.Cache.Enabled {
		if err := s.cache.StoreItems(allItems); err != nil {
			fmt.Fprintf(os.Stderr, "⚠️  Cache storage warning: %v\n", err)
		} else {
			fmt.Fprintf(os.Stderr, "💾 Results cached for future runs\n")
		}
	}

	fmt.Fprintf(os.Stderr, "Discovery completed: %d items\n", len(allItems))

	return allItems, nil
}

func (s *SystemScanner) executeDiscoveryTasks(ctx context.Context, tasks []DiscoveryTask) ([]models.Item, error) {
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, s.safetyLimits.MaxConcurrentScans)
	itemsChan := make(chan []models.Item, len(tasks))
	errorsChan := make(chan error, len(tasks))

	for _, task := range tasks {
		if !task.Enabled {
			continue
		}

		wg.Add(1)
		go s.executeTaskSafely(&wg, semaphore, ctx, itemsChan, errorsChan, task)
	}

	go func() {
		wg.Wait()
		close(itemsChan)
		close(errorsChan)
	}()

	return s.collectResults(itemsChan, errorsChan)
}

func (s *SystemScanner) executeTaskSafely(
	wg *sync.WaitGroup,
	semaphore chan struct{},
	ctx context.Context,
	itemsChan chan<- []models.Item,
	errorsChan chan<- error,
	task DiscoveryTask,
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
				errorsChan <- fmt.Errorf("task %s panicked: %v", task.Name, r)
			}
		}()

		items, err := task.Execute(ctx)
		if err != nil {
			errorsChan <- fmt.Errorf("task %s failed: %w", task.Name, err)
		} else {
			itemsChan <- items
		}
	}()
}

func (s *SystemScanner) collectResults(itemsChan <-chan []models.Item, errorsChan <-chan error) ([]models.Item, error) {
	var allItems []models.Item

	for items := range itemsChan {
		allItems = append(allItems, items...)

		if len(allItems) > s.config.Discovery.MaxExecutables*5 {
			fmt.Fprintf(os.Stderr, "Safety limit reached, truncating results\n")
			break
		}
	}

	for err := range errorsChan {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}

	return allItems, nil
}

func (s *SystemScanner) scanConfiguredPaths(ctx context.Context) ([]models.Item, error) {
	paths := s.config.Discovery.GetEffectivePaths()
	validPaths := s.filterAccessiblePaths(paths)

	if len(validPaths) == 0 {
		return []models.Item{}, nil
	}

	fmt.Fprintf(os.Stderr, "Scanning %d configured paths...\n", len(validPaths))

	return s.scanPathsConcurrently(ctx, validPaths, models.SystemCommand)
}

func (s *SystemScanner) scanPathsConcurrently(ctx context.Context, paths []string, itemType models.ItemType) ([]models.Item, error) {
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, s.safetyLimits.MaxConcurrentScans)
	itemsChan := make(chan []models.Item, len(paths))
	errorsChan := make(chan error, len(paths))

	for _, path := range paths {
		wg.Add(1)
		go func(dirPath string) {
			defer wg.Done()

			select {
			case semaphore <- struct{}{}:
				defer func() { <-semaphore }()
			case <-ctx.Done():
				return
			}

			items, err := s.scanDirectory(dirPath, itemType)
			if err != nil {
				errorsChan <- err
				return
			}

			if len(items) > 0 {
				fmt.Fprintf(os.Stderr, "  %s: %d executables\n", dirPath, len(items))
			}
			itemsChan <- items
		}(path)
	}

	go func() {
		wg.Wait()
		close(itemsChan)
		close(errorsChan)
	}()

	var allItems []models.Item
	for items := range itemsChan {
		allItems = append(allItems, items...)
	}

	for err := range errorsChan {
		fmt.Fprintf(os.Stderr, "⚠️  %v\n", err)
	}

	return allItems, nil
}

func (s *SystemScanner) scanShellConfigs(ctx context.Context) ([]models.Item, error) {
	configFiles := s.getShellConfigFiles()
	var allItems []models.Item

	for _, configFile := range configFiles {
		select {
		case <-ctx.Done():
			return allItems, ctx.Err()
		default:
		}

		items, err := s.parseShellConfig(configFile)
		if err != nil {
			continue
		}

		allItems = append(allItems, items...)
	}

	return allItems, nil
}

func (s *SystemScanner) scanDirectory(dirPath string, itemType models.ItemType) ([]models.Item, error) {
	if !s.isPathSafe(dirPath) {
		return []models.Item{}, fmt.Errorf("path access denied: %s", dirPath)
	}

	if !s.isDirectoryAccessible(dirPath) {
		return []models.Item{}, nil
	}

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory %s: %w", dirPath, err)
	}

	return s.processDirectoryEntries(dirPath, entries, itemType)
}

func (s *SystemScanner) processDirectoryEntries(dirPath string, entries []os.DirEntry, itemType models.ItemType) ([]models.Item, error) {
	var items []models.Item
	helpAttemptCount := 0
	processedCount := 0

	for _, entry := range entries {
		if processedCount >= s.safetyLimits.MaxExecutablesPerDir {
			fmt.Fprintf(os.Stderr, "Reached limit for directory: %s\n", dirPath)
			break
		}

		if entry.IsDir() {
			continue
		}

		filename := entry.Name()

		if s.isDangerousExecutable(filename) {
			continue
		}

		if s.shouldExcludeFile(filename) {
			continue
		}

		fullPath := filepath.Join(dirPath, filename)
		if !s.isExecutable(fullPath) {
			continue
		}

		description := s.generateDescription(filename, fullPath, &helpAttemptCount)

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

func (s *SystemScanner) generateDescription(cmdName, cmdPath string, helpAttemptCount *int) string {
	if s.shouldSkipHelpAttempt(cmdName) || *helpAttemptCount >= s.safetyLimits.MaxHelpAttempts {
		return s.inferDescriptionFromName(cmdName)
	}

	if desc := s.extractHelpDescription(cmdPath); desc != "" {
		*helpAttemptCount++
		return desc
	}

	if desc := s.getManPageDescription(cmdName); desc != "" {
		return desc
	}

	return s.inferDescriptionFromName(cmdName)
}

func (s *SystemScanner) extractHelpDescription(cmdPath string) string {
	filename := filepath.Base(cmdPath)
	if s.isDangerousExecutable(filename) || s.shouldSkipHelpAttempt(filename) {
		return ""
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.safetyLimits.HelpCommandTimeout)
	defer cancel()

	for _, flag := range []string{"--help", "-h"} {
		select {
		case <-ctx.Done():
			return ""
		default:
		}

		if desc := s.tryHelpFlag(ctx, cmdPath, flag); desc != "" {
			return desc
		}
	}

	return ""
}

func (s *SystemScanner) tryHelpFlag(ctx context.Context, cmdPath, flag string) string {
	cmd := exec.CommandContext(ctx, cmdPath, flag)
	cmd.Env = []string{"PATH="}

	output, err := cmd.CombinedOutput()
	if err != nil {
		return ""
	}

	return s.parseHelpOutput(string(output))
}

func (s *SystemScanner) parseHelpOutput(output string) string {
	if len(output) > 10000 {
		output = output[:10000]
	}

	lines := strings.Split(output, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if s.isValidDescriptionLine(line) {
			return line
		}
	}

	return ""
}

func (s *SystemScanner) isValidDescriptionLine(line string) bool {
	if len(line) < 10 || len(line) > 200 {
		return false
	}

	lowerLine := strings.ToLower(line)
	excludePatterns := []string{"usage:", "error:", "fatal:", "warning:", "note:"}

	for _, pattern := range excludePatterns {
		if strings.Contains(lowerLine, pattern) {
			return false
		}
	}

	return true
}

func (s *SystemScanner) scanLegacyPATH(ctx context.Context) ([]models.Item, error) {
	fmt.Fprintf(os.Stderr, "DEPRECATED: Legacy PATH scanning enabled\n")

	if runtime.GOOS == "windows" && s.config.Discovery.WindowsSafeMode {
		fmt.Fprintf(os.Stderr, "Windows Safe Mode: Skipping PATH scan\n")
		return []models.Item{}, nil
	}

	pathEnv := os.Getenv("PATH")
	if pathEnv == "" {
		return []models.Item{}, nil
	}

	separator := ":"
	if runtime.GOOS == "windows" {
		separator = ";"
	}

	paths := strings.Split(pathEnv, separator)
	safePaths := s.filterAccessiblePaths(paths)

	return s.scanPathsConcurrently(ctx, safePaths, models.SystemCommand)
}

func (s *SystemScanner) scanLegacyCommonPaths(ctx context.Context) ([]models.Item, error) {
	fmt.Fprintf(os.Stderr, "DEPRECATED: Legacy common paths scanning enabled\n")

	if runtime.GOOS == "windows" && s.config.Discovery.WindowsSafeMode {
		fmt.Fprintf(os.Stderr, "Windows Safe Mode: Skipping common paths scan\n")
		return []models.Item{}, nil
	}

	commonPaths := s.getLegacyCommonPaths()
	safePaths := s.filterAccessiblePaths(commonPaths)

	return s.scanPathsConcurrently(ctx, safePaths, models.SystemCommand)
}

func (s *SystemScanner) filterAccessiblePaths(paths []string) []string {
	seen := make(map[string]bool)
	var validPaths []string

	for _, path := range paths {
		if path == "" {
			continue
		}

		expandedPath := s.expandPath(path)
		normalizedPath := strings.ToLower(filepath.Clean(expandedPath))

		if seen[normalizedPath] {
			continue
		}
		seen[normalizedPath] = true

		if !s.isPathSafe(expandedPath) {
			continue
		}

		if s.isDirectoryAccessible(expandedPath) {
			validPaths = append(validPaths, expandedPath)
		}
	}

	return validPaths
}

func (s *SystemScanner) isPathSafe(dirPath string) bool {
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

func (s *SystemScanner) expandPath(path string) string {
	if path == "~" || strings.HasPrefix(path, "~/") {
		if homeDir, err := os.UserHomeDir(); err == nil {
			if path == "~" {
				return homeDir
			}
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

func (s *SystemScanner) isExecutable(filePath string) bool {
	info, err := os.Stat(filePath)
	if err != nil {
		return false
	}

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

	return info.Mode()&0111 != 0
}

func (s *SystemScanner) shouldExcludeFile(filename string) bool {
	for _, pattern := range s.config.Discovery.ExcludePatterns {
		if matched, _ := filepath.Match(pattern, filename); matched {
			return true
		}
	}
	return false
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

	skipPatterns := []string{
		"apphostname", "systemsettings", "useraccount", "windowsaction",
		"adminflows", "removedevice", "registrationverifier",
		"shutdown", "format", "diskpart", "regedit", "powercfg",
		"bcdedit", "dism", "sfc", "chkdsk", "msiexec",
		"setup", "install", "uninstall", "config", "admin",
		"control", "manage", "microsoft", "windows",
		".dll", ".sys", ".msi", ".bat", ".cmd", ".ps1", ".vbs", ".js",
	}

	for _, pattern := range skipPatterns {
		if strings.Contains(lower, pattern) {
			return true
		}
	}

	return false
}

func (s *SystemScanner) getShellConfigFiles() []string {
	configFiles := make([]string, 0, len(s.config.Discovery.CustomConfigFiles)+10)
	configFiles = append(configFiles, s.config.Discovery.CustomConfigFiles...)
	configFiles = append(configFiles, s.getDefaultConfigFiles()...)
	return configFiles
}

func (s *SystemScanner) getDefaultConfigFiles() []string {
	homeDir, _ := os.UserHomeDir()

	if runtime.GOOS == "windows" {
		return []string{
			filepath.Join(homeDir, "Documents", "PowerShell", "profile.ps1"),
			filepath.Join(homeDir, ".gitconfig"),
		}
	}

	return []string{
		filepath.Join(homeDir, ".bashrc"),
		filepath.Join(homeDir, ".bash_profile"),
		filepath.Join(homeDir, ".zshrc"),
		filepath.Join(homeDir, ".profile"),
		filepath.Join(homeDir, ".bash_aliases"),
		filepath.Join(homeDir, ".config", "fish", "config.fish"),
		"/etc/bash.bashrc",
		"/etc/bashrc",
		"/etc/profile",
		"/etc/zsh/zshrc",
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

	patterns := struct {
		alias        *regexp.Regexp
		function     *regexp.Regexp
		bashFunction *regexp.Regexp
	}{
		alias:        regexp.MustCompile(`^\s*alias\s+([^=\s]+)=(.*)$`),
		function:     regexp.MustCompile(`^\s*function\s+([^\s\(]+)`),
		bashFunction: regexp.MustCompile(`^\s*([a-zA-Z_][a-zA-Z0-9_]*)\s*\(\s*\)\s*\{`),
	}

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		items = append(items, s.parseConfigLine(line, configPath, patterns)...)
	}

	return items, scanner.Err()
}

func (s *SystemScanner) parseConfigLine(line, configPath string, patterns struct {
	alias        *regexp.Regexp
	function     *regexp.Regexp
	bashFunction *regexp.Regexp
}) []models.Item {
	var items []models.Item

	if matches := patterns.alias.FindStringSubmatch(line); matches != nil {
		aliasName := matches[1]
		aliasValue := strings.Trim(matches[2], `"'`)

		items = append(items, models.Item{
			Name:        aliasName,
			Description: fmt.Sprintf("Alias for: %s", aliasValue),
			Type:        models.UserAlias,
			Source:      configPath,
		})
	}

	if matches := patterns.function.FindStringSubmatch(line); matches != nil {
		functionName := matches[1]
		items = append(items, models.Item{
			Name:        functionName,
			Description: "User-defined shell function",
			Type:        models.UserFunction,
			Source:      configPath,
		})
	}

	if matches := patterns.bashFunction.FindStringSubmatch(line); matches != nil {
		functionName := matches[1]
		items = append(items, models.Item{
			Name:        functionName,
			Description: fmt.Sprintf("Shell function: %s", functionName),
			Type:        models.UserFunction,
			Source:      configPath,
		})
	}

	return items
}

func (s *SystemScanner) getManPageDescription(cmdName string) string {
	if runtime.GOOS == "windows" {
		return ""
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "man", "-f", cmdName)
	output, err := cmd.Output()
	if err != nil {
		return ""
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.Contains(line, " - ") {
			parts := strings.SplitN(line, " - ", 2)
			if len(parts) == 2 {
				return strings.TrimSpace(parts[1])
			}
		}
	}

	return ""
}

func (s *SystemScanner) inferDescriptionFromName(cmdName string) string {
	namePatterns := map[string]string{
		"git":     "Version control system",
		"docker":  "Container management platform",
		"kubectl": "Kubernetes command-line tool",
		"npm":     "Node.js package manager",
		"pip":     "Python package installer",
		"cargo":   "Rust package manager",
		"go":      "Go programming language tool",
		"python":  "Python interpreter",
		"node":    "Node.js JavaScript runtime",
		"java":    "Java application launcher",
		"gcc":     "GNU Compiler Collection",
		"clang":   "C language family frontend",
		"make":    "Build automation tool",
		"cmake":   "Cross-platform build system",
		"tldr":    "Simplified command examples",
		"bat":     "Enhanced cat with syntax highlighting",
		"fd":      "Fast alternative to find",
		"rg":      "Fast text search tool (ripgrep)",
		"exa":     "Modern ls replacement",
		"ls":      "List directory contents",
		"cat":     "Display file contents",
		"grep":    "Search text patterns",
		"find":    "Search for files and directories",
		"curl":    "Transfer data from servers",
		"wget":    "Download files from web",
		"ssh":     "Secure shell remote access",
		"scp":     "Secure copy over network",
		"tar":     "Archive and compress files",
		"zip":     "Create compressed archives",
		"unzip":   "Extract compressed archives",
	}

	lowerName := strings.ToLower(cmdName)
	for pattern, desc := range namePatterns {
		if strings.Contains(lowerName, pattern) {
			return desc
		}
	}

	return "System command"
}

func (s *SystemScanner) postProcessItems(items []models.Item) []models.Item {
	items = s.deduplicateItems(items)

	if len(items) > s.config.Discovery.MaxExecutables {
		fmt.Fprintf(os.Stderr, "Limiting to %d executables (found %d)\n",
			s.config.Discovery.MaxExecutables, len(items))
		items = items[:s.config.Discovery.MaxExecutables]
	}

	return items
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

func (s *SystemScanner) ClearCache() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	fmt.Fprintf(os.Stderr, "Clearing discovery cache...\n")
	return s.cache.ClearCache()
}

// Legacy compatibility method
func (s *SystemScanner) getLegacyCommonPaths() []string {
	switch runtime.GOOS {
	case "windows":
		return []string{
			`C:\ProgramData\chocolatey\bin`,
		}
	case "darwin":
		return []string{
			"/usr/local/bin",
			"/opt/homebrew/bin",
			"/usr/local/sbin",
			"/opt/local/bin",
		}
	default:
		return []string{
			"/usr/local/bin",
			"/usr/bin",
			"/bin",
			"/snap/bin",
			"/var/lib/flatpak/exports/bin",
			"/opt/bin",
			"/usr/local/sbin",
			"/usr/sbin",
			"/sbin",
		}
	}
}
