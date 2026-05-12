package restorers

import "github.com/inuoluwadunsimi/backpack/internal/snapshot"

type PipRestorer struct{}

func (p *PipRestorer) Name() string { return "pip" }

func (p *PipRestorer) Restore(manifest *snapshot.ToolsManifest, dryRun bool) error {
	if manifest.PipPackages == nil {
		return nil
	}
	// TODO: pip3 install <package>==<version> for each package
	return nil
}
