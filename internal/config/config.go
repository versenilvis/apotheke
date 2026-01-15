package config

import (
	"os"
	"path/filepath"
)

type Config struct {
	DataDir string
	DBPath  string
}

func DefaultConfig() (*Config, error) {
	dataDir, err := getDataDir()
	if err != nil {
		return nil, err
	}

	return &Config{
		DataDir: dataDir,
		DBPath:  filepath.Join(dataDir, "apotheke.db"),
	}, nil
}

func getDataDir() (string, error) {
	if xdgData := os.Getenv("XDG_DATA_HOME"); xdgData != "" {
		dir := filepath.Join(xdgData, "apotheke")
		return ensureDir(dir)
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	dir := filepath.Join(home, ".local", "share", "apotheke")
	return ensureDir(dir)
}

func ensureDir(dir string) (string, error) {
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}
	return dir, nil
}
