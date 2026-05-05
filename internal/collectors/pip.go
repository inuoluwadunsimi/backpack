package collectors

import "github.com/inuoluwadunsimi/backpack/internal/snapshot"

// PipCollector captures globally installed pip/pip3 packages.
type PipCollector struct{}

func (p *PipCollector) Name() string { return "pip" }

func (p *PipCollector) IsAvailable() bool {
	// TODO: check if `pip3` is in PATH
	return false
}

func (p *PipCollector) Collect() (*snapshot.ToolState, error) {
	// TODO: run `pip3 list --format=json`
	return &snapshot.ToolState{Name: p.Name()}, nil
}
