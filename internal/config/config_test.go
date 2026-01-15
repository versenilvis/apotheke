package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDefaultConfig(t *testing.T) {
	cfg, err := DefaultConfig()
	if err != nil {
		t.Fatalf("DefaultConfig() error = %v", err)
	}

	if cfg.DataDir == "" {
		t.Error("DataDir should not be empty")
	}

	if cfg.DBPath == "" {
		t.Error("DBPath should not be empty")
	}

	if !filepath.IsAbs(cfg.DBPath) {
		t.Errorf("DBPath should be absolute, got %q", cfg.DBPath)
	}

	expectedDBName := "apotheke.db"
	if filepath.Base(cfg.DBPath) != expectedDBName {
		t.Errorf("DBPath should end with %q, got %q", expectedDBName, cfg.DBPath)
	}
}

func TestDefaultConfig_XDGDataHome(t *testing.T) {
	tmpDir := t.TempDir()
	os.Setenv("XDG_DATA_HOME", tmpDir)
	defer os.Unsetenv("XDG_DATA_HOME")

	cfg, err := DefaultConfig()
	if err != nil {
		t.Fatalf("DefaultConfig() error = %v", err)
	}

	expectedDir := filepath.Join(tmpDir, "apotheke")
	if cfg.DataDir != expectedDir {
		t.Errorf("DataDir = %q, want %q", cfg.DataDir, expectedDir)
	}
}

func TestEnsureDir(t *testing.T) {
	tmpDir := t.TempDir()
	testDir := filepath.Join(tmpDir, "test", "nested", "dir")

	got, err := ensureDir(testDir)
	if err != nil {
		t.Fatalf("ensureDir() error = %v", err)
	}

	if got != testDir {
		t.Errorf("ensureDir() = %q, want %q", got, testDir)
	}

	if _, err := os.Stat(testDir); os.IsNotExist(err) {
		t.Error("directory was not created")
	}
}
