package discovery

import (
	"context"
	"fmt"
	"os"
	"time"
	"what-cmd/internal/config"
	"what-cmd/internal/models"
)

type EmergencyConfig struct {
	DisableAllDiscovery  bool
	DisableHelpExecution bool
	WhitelistOnly        bool
	AllowedExecutables   []string
	MaxDiscoveryTime     time.Duration
	EmergencyShutdown    bool
}

type EmergencyScanner struct {
	config       EmergencyConfig
	innerScanner *SystemScanner
}

func NewEmergencyScanner(cfg *config.Config) (*EmergencyScanner, error) {
	emergencyConfig := EmergencyConfig{
		DisableAllDiscovery:  true,
		DisableHelpExecution: true,
		WhitelistOnly:        true,
		AllowedExecutables: []string{
			"git.exe",
			"node.exe",
			"npm.exe",
			"python.exe",
			"go.exe",
		},
		MaxDiscoveryTime:  5 * time.Second,
		EmergencyShutdown: false,
	}

	scanner := &EmergencyScanner{
		config: emergencyConfig,
	}

	if !emergencyConfig.DisableAllDiscovery {
		innerScanner, err := NewSystemScanner(cfg)
		if err != nil {
			return nil, fmt.Errorf("failed to create inner scanner: %w", err)
		}
		scanner.innerScanner = innerScanner
	}

	return scanner, nil
}

func (es *EmergencyScanner) DiscoverItems(ctx context.Context) ([]models.Item, error) {
	if es.config.EmergencyShutdown {
		return nil, fmt.Errorf("emergency shutdown activated - discovery disabled")
	}

	if es.config.DisableAllDiscovery {
		fmt.Fprintf(os.Stderr, "SAFETY: System discovery disabled for security\n")
		return []models.Item{}, nil
	}

	discoveryCtx, cancel := context.WithTimeout(ctx, es.config.MaxDiscoveryTime)
	defer cancel()

	if es.innerScanner == nil {
		return []models.Item{}, nil
	}

	resultChan := make(chan []models.Item, 1)
	errorChan := make(chan error, 1)

	go func() {
		items, err := es.innerScanner.DiscoverItems(discoveryCtx)
		if err != nil {
			errorChan <- err
			return
		}

		if es.config.WhitelistOnly {
			items = es.filterWhitelistOnly(items)
		}

		resultChan <- items
	}()

	select {
	case items := <-resultChan:
		return items, nil
	case err := <-errorChan:
		return nil, fmt.Errorf("discovery failed safely: %w", err)
	case <-discoveryCtx.Done():
		fmt.Fprintf(os.Stderr, "SAFETY: Discovery timeout - stopping for security\n")
		return []models.Item{}, nil
	}
}

func (es *EmergencyScanner) filterWhitelistOnly(items []models.Item) []models.Item {
	var safe []models.Item

	for _, item := range items {
		if es.isWhitelisted(item.Name) {
			safe = append(safe, item)
		}
	}

	return safe
}

func (es *EmergencyScanner) isWhitelisted(name string) bool {
	for _, allowed := range es.config.AllowedExecutables {
		if name == allowed {
			return true
		}
	}
	return false
}

func (es *EmergencyScanner) ActivateEmergencyShutdown() {
	es.config.EmergencyShutdown = true
	fmt.Fprintf(os.Stderr, "EMERGENCY: System discovery shutdown activated\n")
}
