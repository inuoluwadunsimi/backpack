package snapshot

import "time"

// Snapshot represents a full point-in-time capture of a device's dev environment.
type Snapshot struct {
	ID        string            `json:"id"`
	DeviceID  string            `json:"device_id"`
	Timestamp time.Time         `json:"timestamp"`
	Tools     map[string]ToolState `json:"tools"` // keyed by collector name (e.g. "homebrew", "npm")
	Meta      SnapshotMeta      `json:"meta"`
}

// SnapshotMeta holds metadata about how/why this snapshot was taken.
type SnapshotMeta struct {
	Trigger   string `json:"trigger"`   // "manual", "auto", "init"
	Notes     string `json:"notes"`     // optional user-supplied notes
	Version   string `json:"version"`   // backpack version that created this snapshot
}

// Device represents a registered machine.
type Device struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Hostname string    `json:"hostname"`
	OS       string    `json:"os"`
	Arch     string    `json:"arch"`
	Created  time.Time `json:"created"`
	LastSeen time.Time `json:"last_seen"`
}

// ToolState captures the state of a single tool/package manager.
type ToolState struct {
	Name     string    `json:"name"`
	Version  string    `json:"version,omitempty"`  // version of the tool itself (e.g. brew 4.x)
	Packages []Package `json:"packages"`
	Configs  []Config  `json:"configs,omitempty"`  // relevant config files
}

// Package represents an individual installed package/formula/extension.
type Package struct {
	Name    string `json:"name"`
	Version string `json:"version,omitempty"`
	Source  string `json:"source,omitempty"` // e.g. "tap/formula" for brew, registry URL for npm
	Global  bool   `json:"global,omitempty"` // for npm/pip: globally installed
}

// Config represents a dotfile or config file associated with a tool.
type Config struct {
	Path    string `json:"path"`     // original path on disk (e.g. ~/.zshrc)
	Content string `json:"content"`  // base64-encoded file content
	Hash    string `json:"hash"`     // SHA-256 of raw content for diffing
}

// DiffResult represents the difference between two snapshots or a snapshot and the current state.
type DiffResult struct {
	Tool     string    `json:"tool"`
	Added    []Package `json:"added,omitempty"`
	Removed  []Package `json:"removed,omitempty"`
	Changed  []PackageDiff `json:"changed,omitempty"`
}

// PackageDiff captures a version change for a single package.
type PackageDiff struct {
	Name       string `json:"name"`
	OldVersion string `json:"old_version"`
	NewVersion string `json:"new_version"`
}
