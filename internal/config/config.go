// Package config handles backpack's application configuration.
// Config is stored at ~/.backpack/config.json by default.
package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

const (
	AppDir     = ".backpack"
	ConfigFile = "config.json"
)

// Backend selects the snapshot storage backend.
type Backend string

const (
	BackendLocal Backend = "local"
	BackendGit   Backend = "git"
	BackendS3    Backend = "s3"
)

// Config holds the top-level application configuration.
// Stored at ~/.backpack/config.json.
type Config struct {
	Version      string             `json:"version"`
	Device       DeviceConfig       `json:"device"`
	Storage      StorageConfig      `json:"storage"`
	AutoSnapshot AutoSnapshotConfig `json:"auto_snapshot"`
	Retention    *RetentionConfig   `json:"retention,omitempty"`
}

// DeviceConfig identifies this machine.
type DeviceConfig struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	RegisteredAt time.Time `json:"registered_at"`
}

// StorageConfig selects and configures the snapshot storage backend.
// Only the sub-struct matching Backend should be populated.
type StorageConfig struct {
	Backend Backend             `json:"backend"`
	Local   *LocalStorageConfig `json:"local,omitempty"`
	Git     *GitStorageConfig   `json:"git,omitempty"`
	S3      *S3StorageConfig    `json:"s3,omitempty"`
}

// LocalStorageConfig configures the local filesystem backend.
type LocalStorageConfig struct {
	Path string `json:"path"`
}

// GitStorageConfig configures the git repository backend.
type GitStorageConfig struct {
	RepoURL    string `json:"repo_url"`
	Branch     string `json:"branch"`
	SSHKeyPath string `json:"ssh_key_path,omitempty"`
}

// S3StorageConfig configures the AWS S3 backend.
type S3StorageConfig struct {
	Bucket  string `json:"bucket"`
	Region  string `json:"region"`
	Prefix  string `json:"prefix,omitempty"`
	Profile string `json:"profile,omitempty"`
}

// AutoSnapshotConfig controls automatic snapshot triggers.
type AutoSnapshotConfig struct {
	Enabled bool     `json:"enabled"`
	Hooks   []string `json:"hooks,omitempty"` // tool names that trigger snapshot on change
}

// RetentionConfig is an optional policy for pruning old snapshots.
type RetentionConfig struct {
	MaxSnapshots  int `json:"max_snapshots,omitempty"`  // 0 = unlimited
	RetentionDays int `json:"retention_days,omitempty"` // 0 = forever
}

// LocalPath returns the resolved local storage path from the config.
// Falls back to ~/.backpack/snapshots if no local config is set.
func (c *Config) LocalPath() string {
	if c.Storage.Local != nil && c.Storage.Local.Path != "" {
		return c.Storage.Local.Path
	}
	home, _ := os.UserHomeDir()
	return filepath.Join(home, AppDir, "snapshots")
}

// DefaultConfig returns a sensible default configuration.
func DefaultConfig() *Config {
	home, _ := os.UserHomeDir()
	return &Config{
		Version: "1",
		Storage: StorageConfig{
			Backend: BackendLocal,
			Local: &LocalStorageConfig{
				Path: filepath.Join(home, AppDir, "snapshots"),
			},
		},
		AutoSnapshot: AutoSnapshotConfig{
			Enabled: false,
		},
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
