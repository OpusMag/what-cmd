package search

import (
	"sort"
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
	lowerQuery := strings.ToLower(query)
	sort.Slice(items, func(i, j int) bool {
		si := m.scoreItem(items[i], lowerQuery)
		sj := m.scoreItem(items[j], lowerQuery)
		if si != sj {
			return si > sj
		}
		return items[i].Name < items[j].Name
	})
	return items
}

func (m *Matcher) scoreItem(item models.Item, lowerQuery string) int {
	score := 0
	lowerName := strings.ToLower(item.Name)
	lowerDesc := strings.ToLower(item.Description)

	if strings.EqualFold(item.Name, lowerQuery) {
		score += 100
	} else {

		if strings.Contains(lowerDesc, lowerQuery) {
			score += 10
		}
		if strings.Contains(lowerName, lowerQuery) {
			score += 5
		}

		nameDistance := levenshteinDistance(lowerName, lowerQuery)
		score += max(0, 10-nameDistance)
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


