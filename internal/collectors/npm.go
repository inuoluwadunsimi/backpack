package collectors

import "github.com/inuoluwadunsimi/backpack/internal/snapshot"

// NPMCollector captures globally installed npm packages.
type NPMCollector struct{}

func (n *NPMCollector) Name() string { return "npm" }

func (n *NPMCollector) IsAvailable() bool {
	// TODO: check if `npm` is in PATH
	return false
}

func (n *NPMCollector) Collect() (*snapshot.ToolState, error) {
	// TODO: run `npm list -g --json --depth=0`
	return &snapshot.ToolState{Name: n.Name()}, nil
}
