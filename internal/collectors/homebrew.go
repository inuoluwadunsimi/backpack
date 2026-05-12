package collectors

import (
	"github.com/inuoluwadunsimi/backpack/internal/snapshot"
	"github.com/inuoluwadunsimi/backpack/internal/storage"
)

// HomebrewCollector captures Homebrew formulae, casks, and taps.
type HomebrewCollector struct{}

func (h *HomebrewCollector) Name() string { return "homebrew" }

func (h *HomebrewCollector) IsAvailable() bool {
	// TODO: check if `brew` is in PATH
	return false
}

func (h *HomebrewCollector) Collect(manifest *snapshot.ToolsManifest, _ storage.BlobStore) error {
	// TODO: run `brew --version` to get version
	// TODO: run `brew list --formula --versions` and `brew list --cask --versions`
	// TODO: run `brew tap` to get taps
	// TODO: run `brew list --pinned` to detect pinned formulae
	manifest.Homebrew = &snapshot.HomebrewState{
		Packages: []snapshot.HomebrewPackage{},
		Taps:     []string{},
	}
	return nil
}
