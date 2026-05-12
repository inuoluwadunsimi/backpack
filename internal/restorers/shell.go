package restorers

import (
	"context"

	"github.com/inuoluwadunsimi/backpack/internal/exec"
	"github.com/inuoluwadunsimi/backpack/internal/snapshot"
)

type ShellRestorer struct {
	runner exec.Runner
}

func NewShellRestorer(runner exec.Runner) *ShellRestorer {
	return &ShellRestorer{runner: runner}
}

func (s *ShellRestorer) Name() string { return "shell" }

func (s *ShellRestorer) DependsOn() []string { return []string{"homebrew"} }

func (s *ShellRestorer) CanRun(_ context.Context) bool {
	// Shell restore is always possible
	return true
}

func (s *ShellRestorer) Restore(ctx context.Context, snap *snapshot.Snapshot, opts RestoreOptions) (*RestoreResult, error) {
	if snap.Tools.Shell == nil {
		return &RestoreResult{}, nil
	}
	// TODO: retrieve config file content from blob store via FileRef.Hash
	// TODO: prompt before overwriting existing files
	// TODO: restore .gitconfig, .ssh/config from respective manifest fields
	return &RestoreResult{}, nil
}
