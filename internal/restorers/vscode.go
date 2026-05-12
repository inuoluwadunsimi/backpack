package restorers

import (
	"context"

	"github.com/inuoluwadunsimi/backpack/internal/exec"
	"github.com/inuoluwadunsimi/backpack/internal/snapshot"
)

type VSCodeRestorer struct {
	runner exec.Runner
}

func NewVSCodeRestorer(runner exec.Runner) *VSCodeRestorer {
	return &VSCodeRestorer{runner: runner}
}

func (v *VSCodeRestorer) Name() string { return "vscode" }

func (v *VSCodeRestorer) DependsOn() []string { return []string{"homebrew"} }

func (v *VSCodeRestorer) CanRun(ctx context.Context) bool {
	_, ok := v.runner.Which("code")
	return ok
}

func (v *VSCodeRestorer) Restore(ctx context.Context, snap *snapshot.Snapshot, opts RestoreOptions) (*RestoreResult, error) {
	if snap.Tools.VSCode == nil {
		return &RestoreResult{}, nil
	}
	// TODO: code --install-extension <ext.ID>@<version> for each extension
	// TODO: check existing extensions first → skip if version matches
	// TODO: restore settings.json and keybindings.json from blob store
	return &RestoreResult{}, nil
}
