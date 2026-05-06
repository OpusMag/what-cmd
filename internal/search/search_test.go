package search

import (
	"testing"

	"what-cmd/internal/models"
)

// ─── Matcher.Search ──────────────────────────────────────────────────────────

func TestMatcher_Search_EmptyQuery(t *testing.T) {
	items := []models.Item{
		{Name: "git", Description: "version control"},
		{Name: "docker", Description: "containerization"},
	}
	m := NewMatcher(items)
	result := m.Search("")
	if len(result) != len(items) {
		t.Errorf("empty query: want %d items, got %d", len(items), len(result))
	}
}

func TestMatcher_Search_ExactMatch(t *testing.T) {
	items := []models.Item{
		{Name: "git-status", Description: "show status"},
		{Name: "git", Description: "version control"},
		{Name: "git-log", Description: "show log"},
	}
	m := NewMatcher(items)
	result := m.Search("git")
	if len(result) == 0 {
		t.Fatal("expected results, got none")
	}
	if result[0].Name != "git" {
		t.Errorf("exact match 'git' should rank first, got %q", result[0].Name)
	}
}

func TestMatcher_Search_DescriptionMatch(t *testing.T) {
	items := []models.Item{
		{Name: "docker", Description: "container management"},
		{Name: "unrelated", Description: "something about containers"},
		{Name: "noop", Description: "nothing relevant"},
	}
	m := NewMatcher(items)
	result := m.Search("container")
	if len(result) != 2 {
		t.Errorf("description match: want 2 results, got %d", len(result))
	}
}

func TestMatcher_Search_NoResults(t *testing.T) {
	items := []models.Item{
		{Name: "git", Description: "version control"},
	}
	m := NewMatcher(items)
	result := m.Search("zzzznonexistent")
	if len(result) != 0 {
		t.Errorf("no-match query: want 0 results, got %d", len(result))
	}
}

// ─── Matcher.rankResults ─────────────────────────────────────────────────────

func TestMatcher_rankResults_ExactMatchFirst(t *testing.T) {
	items := []models.Item{
		{Name: "git-status", Description: ""},
		{Name: "git", Description: ""},
	}
	m := &Matcher{}
	result := m.rankResults(items, "git")
	if len(result) == 0 {
		t.Fatal("rankResults returned empty slice")
	}
	if result[0].Name != "git" {
		t.Errorf("exact match should rank first, got %q", result[0].Name)
	}
}

// ─── scoreItem ───────────────────────────────────────────────────────────────

func TestScoreItem_ExactMatchBonus(t *testing.T) {
	item := models.Item{Name: "git", Description: "version control"}
	m := &Matcher{}
	score := m.scoreItem(item, "git")
	if score != 100 {
		t.Errorf("exact match: want score 100, got %d", score)
	}
}

func TestScoreItem_NameContainsBonus(t *testing.T) {
	item := models.Item{Name: "git-status", Description: "show git status"}
	m := &Matcher{}
	score := m.scoreItem(item, "git")
	// name contains (+5) + desc contains (+10) = at least 15
	if score < 15 {
		t.Errorf("name+desc contains bonus: want score >= 15, got %d", score)
	}
}

// TestScoreItem_NoDescriptionLevenshtein is a FIX-6 regression test.
// When description exactly equals the query but the name is long and unrelated:
//   - desc contains match  → +10
//   - name does not contain query → +0
//   - Levenshtein("longcommandname","git") ≥ 12 → max(0, 10-12) = 0
//   - Total expected: 10
//
// If FIX-6 were NOT applied the old code added:
//   descDistance = levenshtein("git","git") = 0 → max(0,5-0) = +5 → total 15.
// Score == 10 confirms description Levenshtein has been removed.
func TestScoreItem_NoDescriptionLevenshtein(t *testing.T) {
	item := models.Item{Name: "longcommandname", Description: "git"}
	m := &Matcher{}
	score := m.scoreItem(item, "git")
	if score != 10 {
		t.Errorf("FIX-6 regression: want score 10 (no desc Levenshtein), got %d", score)
	}
}

// ─── levenshteinDistance ─────────────────────────────────────────────────────

func TestLevenshteinDistance(t *testing.T) {
	cases := []struct {
		s1, s2 string
		want   int
	}{
		{"", "", 0},
		{"abc", "", 3},
		{"", "abc", 3},
		{"abc", "abc", 0},
		{"kitten", "sitting", 3},
		{"git", "git", 0},
		{"git", "got", 1},
		{"a", "b", 1},
	}
	for _, tc := range cases {
		got := levenshteinDistance(tc.s1, tc.s2)
		if got != tc.want {
			t.Errorf("levenshteinDistance(%q, %q) = %d, want %d", tc.s1, tc.s2, got, tc.want)
		}
	}
}
