package search

import (
	"sort"
	"what-cmd/internal/models"
)

type SystemAwareMatcher struct {
	*Matcher
	enableGrouping bool
	scoringWeights ScoringWeights
}

type ScoringWeights struct {
	BuiltinItems     int
	UserDefined      int
	SystemCommands   int
	CustomCommands   int
	DescriptionBonus int
}

func DefaultScoringWeights() ScoringWeights {
	return ScoringWeights{
		BuiltinItems:     20,
		UserDefined:      15,
		SystemCommands:   10,
		CustomCommands:   5,
		DescriptionBonus: 5,
	}
}

func NewSystemAwareMatcher(items []models.Item, enableGrouping bool) *SystemAwareMatcher {
	return &SystemAwareMatcher{
		Matcher:        NewMatcher(items),
		enableGrouping: enableGrouping,
		scoringWeights: DefaultScoringWeights(),
	}
}

func (sm *SystemAwareMatcher) SearchWithGrouping(query string) map[string][]models.Item {
	results := sm.Search(query)

	if !sm.enableGrouping {
		return map[string][]models.Item{"all": results}
	}

	grouped := make(map[string][]models.Item)

	for _, item := range results {
		group := sm.getItemGroup(item)
		grouped[group] = append(grouped[group], item)
	}

	for group := range grouped {
		sm.sortItemsByScore(grouped[group], query)
	}

	return grouped
}

func (sm *SystemAwareMatcher) getItemGroup(item models.Item) string {
	switch item.Type {
	case models.Command, models.Flag, models.Hotkey:
		return "Built-in"
	case models.SystemCommand:
		return "System Commands"
	case models.CustomCommand:
		return "Custom Commands"
	case models.UserAlias:
		return "User Aliases"
	case models.UserFunction:
		return "User Functions"
	case models.PackageManager:
		return "Package Commands"
	default:
		return "Other"
	}
}

func (sm *SystemAwareMatcher) scoreItem(item models.Item, query string) int {
	baseScore := sm.Matcher.scoreItem(item, query)

	switch item.Type {
	case models.Command, models.Flag, models.Hotkey:
		baseScore += sm.scoringWeights.BuiltinItems
	case models.UserAlias, models.UserFunction:
		baseScore += sm.scoringWeights.UserDefined
	case models.SystemCommand:
		baseScore += sm.scoringWeights.SystemCommands
	case models.CustomCommand:
		baseScore += sm.scoringWeights.CustomCommands
	}

	if len(item.Description) > 50 {
		baseScore += sm.scoringWeights.DescriptionBonus
	}

	if item.Source != "" {
		baseScore += 2
	}

	return baseScore
}

func (sm *SystemAwareMatcher) sortItemsByScore(items []models.Item, query string) {
	sort.Slice(items, func(i, j int) bool {
		scoreI := sm.scoreItem(items[i], query)
		scoreJ := sm.scoreItem(items[j], query)

		if scoreI != scoreJ {
			return scoreI > scoreJ
		}

		return items[i].Name < items[j].Name
	})
}

func (sm *SystemAwareMatcher) SetScoringWeights(weights ScoringWeights) {
	sm.scoringWeights = weights
}

func (sm *SystemAwareMatcher) GetGroupedStats() map[string]int {
	stats := make(map[string]int)

	for _, item := range sm.items {
		group := sm.getItemGroup(item)
		stats[group]++
	}

	return stats
}

func (sm *SystemAwareMatcher) FilterByType(itemType models.ItemType) []models.Item {
	var filtered []models.Item

	for _, item := range sm.items {
		if item.Type == itemType {
			filtered = append(filtered, item)
		}
	}

	return filtered
}
