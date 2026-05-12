package collectors

import (
	"context"

	"github.com/inuoluwadunsimi/backpack/internal/exec"
	"github.com/inuoluwadunsimi/backpack/internal/snapshot"
)

// HomebrewCollector captures Homebrew formulae, casks, and taps.
type HomebrewCollector struct {
	runner exec.Runner
}

func NewHomebrewCollector(runner exec.Runner) *HomebrewCollector {
	return &HomebrewCollector{runner: runner}
}

func (h *HomebrewCollector) Name() string { return "homebrew" }

func (h *HomebrewCollector) Available() bool {
	_, ok := h.runner.Which("brew")
	return ok
}

func (h *HomebrewCollector) Collect(ctx context.Context) (*CollectorResult, error) {
	var warnings []string

	// TODO: run `brew --version` to get version
	// TODO: run `brew list --formula --versions` and `brew list --cask --versions`
	//       → parse into []HomebrewPackage with Type "formula"/"cask"
	// TODO: run `brew tap` to get taps
	// TODO: run `brew list --pinned` to detect pinned formulae

	state := &snapshot.HomebrewState{
		Packages: []snapshot.HomebrewPackage{},
		Taps:     []string{},
	}

	return &CollectorResult{
		Available: true,
		Data:      state,
		Warnings:  warnings,
	}, nil
}
