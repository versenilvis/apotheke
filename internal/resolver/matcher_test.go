package resolver

import (
	"testing"
	"time"

	"github.com/versenilvis/apotheke/internal/model"
)

func TestResolver_Resolve_ExactMatch(t *testing.T) {
	r := New()
	now := time.Now()
	commands := []*model.Command{
		{Name: "kdp", Cmd: "kubectl delete pod", Frequency: 10, LastUsed: &now},
		{Name: "kd", Cmd: "kubectl delete", Frequency: 5, LastUsed: &now},
		{Name: "kdps", Cmd: "kubectl delete pods", Frequency: 3, LastUsed: &now},
	}

	matches := r.Resolve("kd", commands)

	if len(matches) != 1 {
		t.Fatalf("Resolve() returned %d matches, want 1 for exact match", len(matches))
	}
	if matches[0].Command.Name != "kd" {
		t.Errorf("Resolve() matched %q, want %q", matches[0].Command.Name, "kd")
	}
}

func TestResolver_Resolve_PrefixMatch(t *testing.T) {
	r := New()
	now := time.Now()
	commands := []*model.Command{
		{Name: "kubectl-delete", Cmd: "kubectl delete", Frequency: 10, LastUsed: &now},
		{Name: "kubectl-get", Cmd: "kubectl get", Frequency: 5, LastUsed: &now},
		{Name: "docker-ps", Cmd: "docker ps", Frequency: 3, LastUsed: &now},
	}

	matches := r.Resolve("kubectl", commands)

	if len(matches) != 2 {
		t.Fatalf("Resolve() returned %d matches, want 2 for prefix match", len(matches))
	}

	for _, m := range matches {
		if m.Command.Name != "kubectl-delete" && m.Command.Name != "kubectl-get" {
			t.Errorf("unexpected match: %q", m.Command.Name)
		}
	}
}

func TestResolver_Resolve_FuzzyMatch(t *testing.T) {
	r := New()
	commands := []*model.Command{
		{Name: "kubectl-delete-pod", Cmd: "kubectl delete pod"},
		{Name: "docker-ps", Cmd: "docker ps"},
	}

	matches := r.Resolve("kdp", commands)

	if len(matches) == 0 {
		t.Fatal("Resolve() returned no matches for fuzzy query")
	}
	if matches[0].Command.Name != "kubectl-delete-pod" {
		t.Errorf("Resolve() matched %q, want %q", matches[0].Command.Name, "kubectl-delete-pod")
	}
}

func TestResolver_Resolve_NoMatch(t *testing.T) {
	r := New()
	commands := []*model.Command{
		{Name: "foo", Cmd: "echo foo"},
		{Name: "bar", Cmd: "echo bar"},
	}

	matches := r.Resolve("xyz", commands)

	if len(matches) != 0 {
		t.Errorf("Resolve() returned %d matches, want 0", len(matches))
	}
}

func TestResolver_Resolve_EmptyCommands(t *testing.T) {
	r := New()
	matches := r.Resolve("test", nil)

	if matches != nil {
		t.Errorf("Resolve() should return nil for empty commands")
	}
}

func TestResolver_Resolve_FrecencyRanking(t *testing.T) {
	r := New()
	recent := time.Now()
	old := time.Now().Add(-30 * 24 * time.Hour)

	commands := []*model.Command{
		{Name: "test-old", Cmd: "echo old", Frequency: 100, LastUsed: &old},
		{Name: "test-recent", Cmd: "echo recent", Frequency: 10, LastUsed: &recent},
	}

	matches := r.Resolve("test", commands)

	if len(matches) != 2 {
		t.Fatalf("Resolve() returned %d matches, want 2", len(matches))
	}

	if matches[0].Command.Name != "test-recent" {
		t.Errorf("recent command should rank higher, got %q first", matches[0].Command.Name)
	}
}

func TestFrecencyScore(t *testing.T) {
	now := time.Now()
	hourAgo := time.Now().Add(-1 * time.Hour)
	dayAgo := time.Now().Add(-24 * time.Hour)
	weekAgo := time.Now().Add(-7 * 24 * time.Hour)

	tests := []struct {
		name     string
		cmd      *model.Command
		minScore int
	}{
		{"no last used", &model.Command{Frequency: 10, LastUsed: nil}, 10},
		{"just used", &model.Command{Frequency: 10, LastUsed: &now}, 90},
		{"hour ago", &model.Command{Frequency: 10, LastUsed: &hourAgo}, 80},
		{"day ago", &model.Command{Frequency: 10, LastUsed: &dayAgo}, 40},
		{"week ago", &model.Command{Frequency: 10, LastUsed: &weekAgo}, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			score := frecencyScore(tt.cmd)
			if score < tt.minScore {
				t.Errorf("frecencyScore() = %d, want >= %d", score, tt.minScore)
			}
		})
	}
}
