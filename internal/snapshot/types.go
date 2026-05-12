package snapshot

import "time"

// ──────────────────────────────────────────────
// Snapshot — top-level
// ──────────────────────────────────────────────

// Snapshot represents a full point-in-time capture of a device's dev environment.
type Snapshot struct {
	SchemaVersion int            `json:"schema_version"`
	ID            string         `json:"id"`
	CapturedAt    time.Time      `json:"captured_at"`
	Trigger       string         `json:"trigger"`             // "manual", "auto", "init"
	ParentID      string         `json:"parent_id,omitempty"` // previous snapshot ID
	Notes         string         `json:"notes,omitempty"`
	Tags          []string       `json:"tags,omitempty"`
	Device        SnapshotDevice `json:"device"`
	Tools         ToolsManifest  `json:"tools"`
}

// SnapshotDevice is the device context embedded in each snapshot.
// Denormalized from config so snapshots are self-contained.
type SnapshotDevice struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	OS        string `json:"os"`
	Arch      string `json:"arch"`
	OSVersion string `json:"os_version"`
}

// ──────────────────────────────────────────────
// Device (persisted separately in device.json)
// ──────────────────────────────────────────────

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

// ──────────────────────────────────────────────
// Tools Manifest — typed per-tool
// ──────────────────────────────────────────────

// ToolsManifest holds every tool's state as a typed field.
// Nil pointer = tool was not collected (unavailable or disabled).
type ToolsManifest struct {
	Homebrew    *HomebrewState    `json:"homebrew,omitempty"`
	Node        *RuntimeState     `json:"node,omitempty"`
	NpmGlobals  []PackageEntry    `json:"npm_globals,omitempty"`
	Python      *RuntimeState     `json:"python,omitempty"`
	PipPackages []PackageEntry    `json:"pip_packages,omitempty"`
	VSCode      *VSCodeState      `json:"vscode,omitempty"`
	Shell       *ShellState       `json:"shell,omitempty"`
	Git         *GitConfigState   `json:"git,omitempty"`
	SSH         *SSHState         `json:"ssh,omitempty"`
	SystemTools []SystemToolEntry `json:"system_tools,omitempty"`
}

// ──────────────────────────────────────────────
// Homebrew
// ──────────────────────────────────────────────

// HomebrewState captures Homebrew formulae, casks, and taps.
type HomebrewState struct {
	Version  string            `json:"version"`
	Packages []HomebrewPackage `json:"packages"`
	Taps     []string          `json:"taps"`
}

// HomebrewPackage represents a formula or cask.
type HomebrewPackage struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Type    string `json:"type"`             // "formula" or "cask"
	Pinned  bool   `json:"pinned,omitempty"` // only relevant for formulae
}

// ──────────────────────────────────────────────
// Language Runtimes (Node, Python, etc.)
// ──────────────────────────────────────────────

// RuntimeState captures a language runtime's version and how it's managed.
type RuntimeState struct {
	Version string `json:"version"`
	Manager string `json:"manager,omitempty"` // "nvm", "pyenv", "system", etc.
}

// PackageEntry is a globally-installed package (npm -g, pip).
type PackageEntry struct {
	Name    string `json:"name"`
	Version string `json:"version,omitempty"`
}

// ──────────────────────────────────────────────
// VS Code
// ──────────────────────────────────────────────

// VSCodeState captures VS Code version and installed extensions.
type VSCodeState struct {
	Version    string            `json:"version"`
	Extensions []VSCodeExtension `json:"extensions"`
}

// VSCodeExtension represents an installed VS Code extension.
type VSCodeExtension struct {
	ID      string `json:"id"`      // publisher.name format
	Version string `json:"version"`
}

// ──────────────────────────────────────────────
// Shell
// ──────────────────────────────────────────────

// ShellState captures the user's shell configuration.
type ShellState struct {
	Type        string             `json:"type"`                    // "zsh", "bash", "fish"
	ConfigFiles map[string]FileRef `json:"config_files,omitempty"` // filename → blob reference
	Aliases     []string           `json:"aliases,omitempty"`       // raw alias lines
}

// FileRef references a config file whose content is stored in the blob store.
type FileRef struct {
	Hash string `json:"hash"`           // SHA-256 of raw content
	Size int64  `json:"size"`           // raw file size in bytes
	Path string `json:"path"`           // original absolute path on disk
	Mode string `json:"mode,omitempty"` // file permission bits, e.g. "0644"
}

// ──────────────────────────────────────────────
// Git Config
// ──────────────────────────────────────────────

// GitConfigState captures the user's git configuration.
type GitConfigState struct {
	Name    string            `json:"name"`
	Email   string            `json:"email"`
	Editor  string            `json:"editor,omitempty"`
	Aliases map[string]string `json:"aliases,omitempty"` // alias → expansion
}

// ──────────────────────────────────────────────
// SSH
// ──────────────────────────────────────────────

// SSHState captures SSH key presence (public keys only) and config existence.
type SSHState struct {
	KeyFiles  []string `json:"key_files"`  // public key filenames only
	HasConfig bool     `json:"has_config"` // true if ~/.ssh/config exists
}

// ──────────────────────────────────────────────
// System Tools (go, rustc, java, etc.)
// ──────────────────────────────────────────────

// SystemToolEntry represents a standalone CLI tool installed on the system.
type SystemToolEntry struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Path    string `json:"path"` // resolved binary path
}

// ──────────────────────────────────────────────
// Diff Types
// ──────────────────────────────────────────────

// DiffResult represents the difference between two snapshots for one tool.
type DiffResult struct {
	Tool          string        `json:"tool"`
	Added         []PackageEntry `json:"added,omitempty"`
	Removed       []PackageEntry `json:"removed,omitempty"`
	Changed       []PackageDiff  `json:"changed,omitempty"`
	ConfigChanges []ConfigDiff   `json:"config_changes,omitempty"`
}

// PackageDiff captures a version change for a single package.
type PackageDiff struct {
	Name       string `json:"name"`
	OldVersion string `json:"old_version"`
	NewVersion string `json:"new_version"`
}

// ConfigDiff tracks changes to a config file between snapshots.
type ConfigDiff struct {
	Path    string `json:"path"`
	OldHash string `json:"old_hash,omitempty"`
	NewHash string `json:"new_hash,omitempty"`
	Status  string `json:"status"` // "added", "removed", "modified"
}
