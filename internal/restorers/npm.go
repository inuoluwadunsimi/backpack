package restorers

import (
	"context"

	"github.com/inuoluwadunsimi/backpack/internal/exec"
	"github.com/inuoluwadunsimi/backpack/internal/snapshot"
)

type NPMRestorer struct {
	runner exec.Runner
}

func NewNPMRestorer(runner exec.Runner) *NPMRestorer {
	return &NPMRestorer{runner: runner}
}

func (n *NPMRestorer) Name() string { return "npm" }

func (n *NPMRestorer) DependsOn() []string { return []string{"homebrew"} }

func (n *NPMRestorer) CanRun(ctx context.Context) bool {
	_, ok := n.runner.Which("npm")
	return ok
}

func (n *NPMRestorer) Restore(ctx context.Context, snap *snapshot.Snapshot, opts RestoreOptions) (*RestoreResult, error) {
	if snap.Tools.NpmGlobals == nil {
		return &RestoreResult{}, nil
	}
	// TODO: npm install -g <package>@<version> for each package
	// TODO: check existing version first → skip if matches
	return &RestoreResult{}, nil
}
