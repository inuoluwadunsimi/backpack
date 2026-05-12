package restorers

import "github.com/inuoluwadunsimi/backpack/internal/snapshot"

type VSCodeRestorer struct{}

func (v *VSCodeRestorer) Name() string { return "vscode" }

func (v *VSCodeRestorer) Restore(manifest *snapshot.ToolsManifest, dryRun bool) error {
	if manifest.VSCode == nil {
		return nil
	}
	// TODO: code --install-extension <ext.ID> for each extension
	// TODO: restore settings.json and keybindings.json from blob store
	return nil
}
