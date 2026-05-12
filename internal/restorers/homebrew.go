package restorers

import (
	"context"

	"github.com/inuoluwadunsimi/backpack/internal/exec"
	"github.com/inuoluwadunsimi/backpack/internal/snapshot"
)

type HomebrewRestorer struct {
	runner exec.Runner
}

func NewHomebrewRestorer(runner exec.Runner) *HomebrewRestorer {
	return &HomebrewRestorer{runner: runner}
}

func (h *HomebrewRestorer) Name() string { return "homebrew" }

func (h *HomebrewRestorer) DependsOn() []string { return []string{"system"} }

func (h *HomebrewRestorer) CanRun(ctx context.Context) bool {
	_, ok := h.runner.Which("brew")
	return ok
}

func (h *HomebrewRestorer) Restore(ctx context.Context, snap *snapshot.Snapshot, opts RestoreOptions) (*RestoreResult, error) {
	if snap.Tools.Homebrew == nil {
		return &RestoreResult{}, nil
	}
	// TODO: install Homebrew itself if missing
	// TODO: add taps
	// TODO: for each package:
	//   - check if already installed at correct version → skip
	//   - if installed at wrong version → upgrade
	//   - if not installed → install
	//   - on failure → record in Failed, continue
	// TODO: separate formula (brew install) from cask (brew install --cask)
	return &RestoreResult{}, nil
}
