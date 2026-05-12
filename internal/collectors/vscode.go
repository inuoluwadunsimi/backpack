package collectors

import (
	"context"

	"github.com/inuoluwadunsimi/backpack/internal/exec"
	"github.com/inuoluwadunsimi/backpack/internal/snapshot"
)

// VSCodeCollector captures VS Code extensions and settings.
type VSCodeCollector struct {
	runner exec.Runner
}

func NewVSCodeCollector(runner exec.Runner) *VSCodeCollector {
	return &VSCodeCollector{runner: runner}
}

func (v *VSCodeCollector) Name() string { return "vscode" }

func (v *VSCodeCollector) Available() bool {
	_, ok := v.runner.Which("code")
	return ok
}

func (v *VSCodeCollector) Collect(ctx context.Context) (*CollectorResult, error) {
	var warnings []string

	// TODO: run `code --version` to get VS Code version
	// TODO: run `code --list-extensions --show-versions`
	// TODO: optionally capture settings.json and keybindings.json via blobs

	state := &snapshot.VSCodeState{
		Extensions: []snapshot.VSCodeExtension{},
	}

	return &CollectorResult{
		Available: true,
		Data:      state,
		Warnings:  warnings,
	}, nil
}
