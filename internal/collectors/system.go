package collectors

import (
	"runtime"

	"github.com/inuoluwadunsimi/backpack/internal/snapshot"
)

// SystemCollector captures macOS system-level information:
// OS version, Xcode CLI tools, default shell, hardware arch.
type SystemCollector struct{}

func (s *SystemCollector) Name() string { return "system" }

func (s *SystemCollector) IsAvailable() bool {
	return runtime.GOOS == "darwin"
}

func (s *SystemCollector) Collect() (*snapshot.ToolState, error) {
	// TODO: capture sw_vers output (macOS version)
	// TODO: check xcode-select --version
	// TODO: capture arch (arm64 vs x86_64)
	return &snapshot.ToolState{Name: s.Name()}, nil
}
