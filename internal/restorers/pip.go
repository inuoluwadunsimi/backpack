package restorers

import "github.com/inuoluwadunsimi/backpack/internal/snapshot"

type PipRestorer struct{}

func (p *PipRestorer) Name() string { return "pip" }

func (p *PipRestorer) Restore(state *snapshot.ToolState, dryRun bool) error {
	// TODO: pip3 install <package>==<version> for each package
	return nil
}
