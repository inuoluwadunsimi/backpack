package restorers

import "github.com/inuoluwadunsimi/backpack/internal/snapshot"

type HomebrewRestorer struct{}

func (h *HomebrewRestorer) Name() string { return "homebrew" }

func (h *HomebrewRestorer) Restore(state *snapshot.ToolState, dryRun bool) error {
	// TODO: install Homebrew itself if missing
	// TODO: add taps
	// TODO: brew install formulae
	// TODO: brew install --cask casks
	return nil
}
