package collectors

import "github.com/inuoluwadunsimi/backpack/internal/snapshot"

// HomebrewCollector captures Homebrew formulae, casks, and taps.
type HomebrewCollector struct{}

func (h *HomebrewCollector) Name() string { return "homebrew" }

func (h *HomebrewCollector) IsAvailable() bool {
	// TODO: check if `brew` is in PATH
	return false
}

func (h *HomebrewCollector) Collect() (*snapshot.ToolState, error) {
	// TODO: run `brew list --formula --versions`, `brew list --cask --versions`, `brew tap`
	return &snapshot.ToolState{Name: h.Name()}, nil
}
