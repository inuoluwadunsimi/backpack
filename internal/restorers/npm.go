package restorers

import "github.com/inuoluwadunsimi/backpack/internal/snapshot"

type NPMRestorer struct{}

func (n *NPMRestorer) Name() string { return "npm" }

func (n *NPMRestorer) Restore(state *snapshot.ToolState, dryRun bool) error {
	// TODO: npm install -g <package>@<version> for each package
	return nil
}
