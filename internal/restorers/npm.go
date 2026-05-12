package restorers

import "github.com/inuoluwadunsimi/backpack/internal/snapshot"

type NPMRestorer struct{}

func (n *NPMRestorer) Name() string { return "npm" }

func (n *NPMRestorer) Restore(manifest *snapshot.ToolsManifest, dryRun bool) error {
	if manifest.NpmGlobals == nil {
		return nil
	}
	// TODO: npm install -g <package>@<version> for each package
	return nil
}
