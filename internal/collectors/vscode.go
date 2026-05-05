package collectors

import "github.com/inuoluwadunsimi/backpack/internal/snapshot"

// VSCodeCollector captures VS Code extensions and settings.
type VSCodeCollector struct{}

func (v *VSCodeCollector) Name() string { return "vscode" }

func (v *VSCodeCollector) IsAvailable() bool {
	// TODO: check if `code` CLI is in PATH
	return false
}

func (v *VSCodeCollector) Collect() (*snapshot.ToolState, error) {
	// TODO: run `code --list-extensions --show-versions`
	// TODO: optionally capture settings.json and keybindings.json
	return &snapshot.ToolState{Name: v.Name()}, nil
}
