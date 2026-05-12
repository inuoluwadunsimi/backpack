package collectors

import (
	"context"
	"runtime"

	"github.com/inuoluwadunsimi/backpack/internal/exec"
	"github.com/inuoluwadunsimi/backpack/internal/snapshot"
)

// SystemCollector captures macOS system-level information and standalone
// CLI tools (go, rustc, java, etc.).
type SystemCollector struct {
	runner exec.Runner
}

func NewSystemCollector(runner exec.Runner) *SystemCollector {
	return &SystemCollector{runner: runner}
}

func (s *SystemCollector) Name() string { return "system" }

func (s *SystemCollector) Available() bool {
	return runtime.GOOS == "darwin"
}

func (s *SystemCollector) Collect(ctx context.Context) (*CollectorResult, error) {
	var warnings []string

	// TODO: capture sw_vers output (macOS version)
	// TODO: check xcode-select --version
	// TODO: detect standalone tools: go version, rustc --version, java -version

	return &CollectorResult{
		Available: true,
		Data:      []snapshot.SystemToolEntry{},
		Warnings:  warnings,
	}, nil
}
