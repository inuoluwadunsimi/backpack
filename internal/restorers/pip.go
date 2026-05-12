package restorers

import (
	"context"

	"github.com/inuoluwadunsimi/backpack/internal/exec"
	"github.com/inuoluwadunsimi/backpack/internal/snapshot"
)

type PipRestorer struct {
	runner exec.Runner
}

func NewPipRestorer(runner exec.Runner) *PipRestorer {
	return &PipRestorer{runner: runner}
}

func (p *PipRestorer) Name() string { return "pip" }

func (p *PipRestorer) DependsOn() []string { return []string{"homebrew"} }

func (p *PipRestorer) CanRun(ctx context.Context) bool {
	_, ok := p.runner.Which("pip3")
	return ok
}

func (p *PipRestorer) Restore(ctx context.Context, snap *snapshot.Snapshot, opts RestoreOptions) (*RestoreResult, error) {
	if snap.Tools.PipPackages == nil {
		return &RestoreResult{}, nil
	}
	// TODO: pip3 install <package>==<version> for each package
	// TODO: check existing version first → skip if matches
	return &RestoreResult{}, nil
}
