package restorers

import (
	"time"

	"github.com/inuoluwadunsimi/backpack/internal/snapshot"

	"context"
)

// RestoreOptions controls restore behaviour.
type RestoreOptions struct {
	DryRun     bool     // print what would happen without making changes
	ToolFilter []string // only restore these tools (empty = all)
	Force      bool     // reinstall even if version matches
}

// RestoreResult summarises what happened during a restore.
type RestoreResult struct {
	Installed []string     // packages successfully installed/upgraded
	Skipped   []string     // packages already at correct version
	Failed    []FailedItem // packages that failed to install
	Duration  time.Duration
}

// FailedItem records a single package install failure.
type FailedItem struct {
	Name    string
	Version string
	Err     error
	Retried bool // true if we retried and still failed
}

// Restorer reinstalls packages for a single tool from a snapshot.
type Restorer interface {
	// Name returns the tool name this restorer handles.
	Name() string

	// DependsOn returns tool names that must be restored before this one.
	// Used by the dependency graph to determine restore wave ordering.
	DependsOn() []string

	// CanRun checks if this restorer's prerequisites are met (e.g. the
	// package manager is installed). Returns false with no error if the
	// tool simply isn't available on this machine.
	CanRun(ctx context.Context) bool

	// Restore installs packages from the snapshot according to opts.
	// Must never abort the wave on a single item failure — record it in
	// RestoreResult.Failed and continue.
	Restore(ctx context.Context, snap *snapshot.Snapshot, opts RestoreOptions) (*RestoreResult, error)
}
