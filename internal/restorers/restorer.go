// Package restorers defines the Restorer interface and a registry for
// reinstalling tools from a snapshot.
package restorers

import "github.com/inuoluwadunsimi/backpack/internal/snapshot"

// Restorer knows how to reinstall packages for a single tool.
type Restorer interface {
	// Name returns the collector name this restorer corresponds to.
	Name() string

	// Restore installs/configures packages from the given ToolsManifest.
	// If dryRun is true, it should only print what it would do.
	Restore(manifest *snapshot.ToolsManifest, dryRun bool) error
}

// Registry returns all known restorers in dependency order.
// Homebrew is first because npm/pip may be installed via brew.
func Registry() []Restorer {
	return []Restorer{
		&HomebrewRestorer{},
		&ShellRestorer{},
		&NPMRestorer{},
		&PipRestorer{},
		&VSCodeRestorer{},
	}
}
