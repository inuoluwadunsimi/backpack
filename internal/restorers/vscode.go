package restorers

import "github.com/inuoluwadunsimi/backpack/internal/snapshot"

type VSCodeRestorer struct{}

func (v *VSCodeRestorer) Name() string { return "vscode" }

func (v *VSCodeRestorer) Restore(state *snapshot.ToolState, dryRun bool) error {
	// TODO: code --install-extension <ext> for each extension
	// TODO: restore settings.json and keybindings.json
	return nil
}
