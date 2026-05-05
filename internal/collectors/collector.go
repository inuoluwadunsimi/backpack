// Package collectors defines the Collector interface and a registry for
// discovering installed dev tools and capturing their state.
package collectors

import "github.com/inuoluwadunsimi/backpack/internal/snapshot"

// Collector captures the current state of a single tool/package manager.
type Collector interface {
	// Name returns a unique identifier for this collector (e.g. "homebrew").
	Name() string

	// IsAvailable returns true if the tool is installed on this machine.
	IsAvailable() bool

	// Collect gathers the current state of the tool.
	Collect() (*snapshot.ToolState, error)
}

// Registry returns all known collectors. During restore the order matters
// (brew before npm, etc.) so they are returned in dependency order.
func Registry() []Collector {
	return []Collector{
		&SystemCollector{},
		&HomebrewCollector{},
		&ShellCollector{},
		&NPMCollector{},
		&PipCollector{},
		&VSCodeCollector{},
	}
}
