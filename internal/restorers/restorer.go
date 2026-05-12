package restorers

import "github.com/inuoluwadunsimi/backpack/internal/exec"

// Registry returns all known restorers in dependency order.
func Registry(runner exec.Runner) []Restorer {
	return []Restorer{
		NewHomebrewRestorer(runner),
		NewShellRestorer(runner),
		NewNPMRestorer(runner),
		NewPipRestorer(runner),
		NewVSCodeRestorer(runner),
	}
}
