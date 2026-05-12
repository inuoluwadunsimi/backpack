package collectors

import (
	"context"

	"github.com/inuoluwadunsimi/backpack/internal/exec"
	"github.com/inuoluwadunsimi/backpack/internal/snapshot"
)

// PipCollector captures globally installed pip/pip3 packages and the Python runtime.
type PipCollector struct {
	runner exec.Runner
}

func NewPipCollector(runner exec.Runner) *PipCollector {
	return &PipCollector{runner: runner}
}

func (p *PipCollector) Name() string { return "pip" }

func (p *PipCollector) Available() bool {
	_, ok := p.runner.Which("pip3")
	return ok
}

// PipCollectorData bundles python runtime and global packages.
type PipCollectorData struct {
	Python      *snapshot.RuntimeState
	PipPackages []snapshot.PackageEntry
}

func (p *PipCollector) Collect(ctx context.Context) (*CollectorResult, error) {
	var warnings []string

	// TODO: run `python3 --version` to get runtime version
	// TODO: detect manager (pyenv, system)
	// TODO: run `pip3 list --format=json`

	data := &PipCollectorData{
		Python:      &snapshot.RuntimeState{},
		PipPackages: []snapshot.PackageEntry{},
	}

	return &CollectorResult{
		Available: true,
		Data:      data,
		Warnings:  warnings,
	}, nil
}
