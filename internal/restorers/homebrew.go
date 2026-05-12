package restorers

import "github.com/inuoluwadunsimi/backpack/internal/snapshot"

type HomebrewRestorer struct{}

func (h *HomebrewRestorer) Name() string { return "homebrew" }

func (h *HomebrewRestorer) Restore(manifest *snapshot.ToolsManifest, dryRun bool) error {
	if manifest.Homebrew == nil {
		return nil
	}
	// TODO: install Homebrew itself if missing
	// TODO: add taps
	// TODO: brew install formulae (packages where Type == "formula")
	// TODO: brew install --cask casks (packages where Type == "cask")
	return nil
}
