// Package config handles backpack's application configuration.
// Config is stored at ~/.backpack/config.json by default.
package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const (
	AppDir     = ".backpack"
	ConfigFile = "config.json"
)

// Config holds the top-level application configuration.
type Config struct {
	DeviceID   string        `json:"device_id"`
	DeviceName string        `json:"device_name"`
	Storage    StorageConfig `json:"storage"`
	AutoWatch  bool          `json:"auto_watch"` // enable fs-watch based auto-snapshots
}

// StorageConfig selects and configures the snapshot storage backend.
type StorageConfig struct {
	Backend string `json:"backend"` // "local", "git", "s3"

	// Local backend
	LocalPath string `json:"local_path,omitempty"`

	// Git backend
	GitRepo   string `json:"git_repo,omitempty"`
	GitBranch string `json:"git_branch,omitempty"`

	// S3 backend
	S3Bucket string `json:"s3_bucket,omitempty"`
	S3Region string `json:"s3_region,omitempty"`
	S3Prefix string `json:"s3_prefix,omitempty"`
}

// DefaultConfig returns a sensible default configuration.
func DefaultConfig() *Config {
	home, _ := os.UserHomeDir()
	return &Config{
		Storage: StorageConfig{
			Backend:   "local",
			LocalPath: filepath.Join(home, AppDir, "snapshots"),
		},
		AutoWatch: false,
	}
}

// AppDataDir returns the path to ~/.backpack, creating it if necessary.
func AppDataDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("cannot determine home directory: %w", err)
	}
	dir := filepath.Join(home, AppDir)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return "", fmt.Errorf("cannot create app directory: %w", err)
	}
	return dir, nil
}

// Load reads the config from disk. Returns DefaultConfig if the file doesn't exist.
func Load() (*Config, error) {
	dir, err := AppDataDir()
	if err != nil {
		return nil, err
	}

	path := filepath.Join(dir, ConfigFile)
	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return DefaultConfig(), nil
	}
	if err != nil {
		return nil, fmt.Errorf("reading config: %w", err)
	}

	cfg := &Config{}
	if err := json.Unmarshal(data, cfg); err != nil {
		return nil, fmt.Errorf("parsing config: %w", err)
	}
	return cfg, nil
}

// Save writes the config to disk.
func Save(cfg *Config) error {
	dir, err := AppDataDir()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("marshalling config: %w", err)
	}

	path := filepath.Join(dir, ConfigFile)
	return os.WriteFile(path, data, 0o644)
}
