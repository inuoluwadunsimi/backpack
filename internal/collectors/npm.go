package collectors

import (
	"context"

	"github.com/inuoluwadunsimi/backpack/internal/exec"
	"github.com/inuoluwadunsimi/backpack/internal/snapshot"
)

// NPMCollector captures globally installed npm packages and the Node.js runtime.
type NPMCollector struct {
	runner exec.Runner
}

func NewNPMCollector(runner exec.Runner) *NPMCollector {
	return &NPMCollector{runner: runner}
}

func (n *NPMCollector) Name() string { return "npm" }

func (n *NPMCollector) Available() bool {
	_, ok := n.runner.Which("npm")
	return ok
}

// NPMCollectorData bundles node runtime and global packages.
type NPMCollectorData struct {
	Node       *snapshot.RuntimeState
	NpmGlobals []snapshot.PackageEntry
}

func (n *NPMCollector) Collect(ctx context.Context) (*CollectorResult, error) {
	var warnings []string

	// TODO: run `node --version` to get runtime version
	// TODO: detect manager (nvm, fnm, system) by checking which shim resolves
	// TODO: run `npm list -g --json --depth=0` for global packages

	data := &NPMCollectorData{
		Node:       &snapshot.RuntimeState{},
		NpmGlobals: []snapshot.PackageEntry{},
	}

	return &CollectorResult{
		Available: true,
		Data:      data,
		Warnings:  warnings,
	}, nil
}
