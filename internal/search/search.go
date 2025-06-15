package search

import (
	"strings"
	"what-cmd/internal/models"
)

type Matcher struct {
	items []models.Item
}

func NewMatcher(items []models.Item) *Matcher {
	return &Matcher{items: items}
}

func (m *Matcher) Search(query string) []models.Item {
	if query == "" {
		return m.items
	}

	var results []models.Item
	lowerQuery := strings.ToLower(query)

	for _, item := range m.items {
		if m.matches(item, lowerQuery) {
			results = append(results, item)
		}
	}

	return m.rankResults(results, query)
}

func (m *Matcher) matches(item models.Item, lowerQuery string) bool {
	lowerName := strings.ToLower(item.Name)
	lowerDesc := strings.ToLower(item.Description)

	return strings.Contains(lowerName, lowerQuery) ||
		strings.Contains(lowerDesc, lowerQuery)
}

func (m *Matcher) rankResults(items []models.Item, query string) []models.Item {
	if len(items) == 0 {
		return items
	}

	bestIdx := 0
	bestScore := m.scoreItem(items[0], query)

	for i := 1; i < len(items); i++ {
		score := m.scoreItem(items[i], query)
		if score > bestScore {
			bestScore = score
			bestIdx = i
		}
	}

	if bestIdx > 0 {
		items[0], items[bestIdx] = items[bestIdx], items[0]
	}

	return items
}

func (m *Matcher) scoreItem(item models.Item, query string) int {
	score := 0
	lowerQuery := strings.ToLower(query)
	lowerName := strings.ToLower(item.Name)
	lowerDesc := strings.ToLower(item.Description)

	// Exact match gets highest score so the user gets what they want
	if strings.EqualFold(item.Name, query) {
		score += 100
	} else {

		if strings.Contains(lowerDesc, lowerQuery) {
			score += 10
		}
		if strings.Contains(lowerName, lowerQuery) {
			score += 5
		}

		nameDistance := levenshteinDistance(lowerName, lowerQuery)
		descDistance := levenshteinDistance(lowerDesc, lowerQuery)
		score += max(0, 10-nameDistance)
		score += max(0, 5-descDistance)
	}

	return score
}

func levenshteinDistance(s1, s2 string) int {
	len1, len2 := len(s1), len(s2)
	if len1 == 0 {
		return len2
	}
	if len2 == 0 {
		return len1
	}

	matrix := make([][]int, len1+1)
	for i := range matrix {
		matrix[i] = make([]int, len2+1)
	}

	for i := 0; i <= len1; i++ {
		matrix[i][0] = i
	}
	for j := 0; j <= len2; j++ {
		matrix[0][j] = j
	}

	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			if s1[i-1] == s2[j-1] {
				matrix[i][j] = matrix[i-1][j-1]
			} else {
				matrix[i][j] = 1 + min(
					matrix[i-1][j],   // deletion
					matrix[i][j-1],   // insertion
					matrix[i-1][j-1], // substitution
				)
			}
		}
	}

	return matrix[len1][len2]
}

func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= c {
		return b
	}
	return c
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
