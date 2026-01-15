package model

import (
	"testing"
)

func TestCommand_HasTag(t *testing.T) {
	tests := []struct {
		name     string
		tags     string
		tag      string
		expected bool
	}{
		{"empty tags", "", "foo", false},
		{"single tag match", "foo", "foo", true},
		{"single tag no match", "foo", "bar", false},
		{"multiple tags match first", "foo,bar,baz", "foo", true},
		{"multiple tags match middle", "foo,bar,baz", "bar", true},
		{"multiple tags match last", "foo,bar,baz", "baz", true},
		{"multiple tags no match", "foo,bar,baz", "qux", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &Command{Tags: tt.tags}
			if got := cmd.HasTag(tt.tag); got != tt.expected {
				t.Errorf("HasTag(%q) = %v, want %v", tt.tag, got, tt.expected)
			}
		})
	}
}

func TestCommand_IsDangerous(t *testing.T) {
	tests := []struct {
		name     string
		tags     string
		expected bool
	}{
		{"no tags", "", false},
		{"danger tag", "danger", true},
		{"danger with others", "k8s,danger,prod", true},
		{"no danger", "k8s,prod", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &Command{Tags: tt.tags}
			if got := cmd.IsDangerous(); got != tt.expected {
				t.Errorf("IsDangerous() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestSplitTags(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{"", nil},
		{"foo", []string{"foo"}},
		{"foo,bar", []string{"foo", "bar"}},
		{"foo,bar,baz", []string{"foo", "bar", "baz"}},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := splitTags(tt.input)
			if len(got) != len(tt.expected) {
				t.Errorf("splitTags(%q) = %v, want %v", tt.input, got, tt.expected)
				return
			}
			for i := range got {
				if got[i] != tt.expected[i] {
					t.Errorf("splitTags(%q)[%d] = %q, want %q", tt.input, i, got[i], tt.expected[i])
				}
			}
		})
	}
}
