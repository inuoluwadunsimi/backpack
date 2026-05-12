package collectors

import (
	"runtime"

	"github.com/inuoluwadunsimi/backpack/internal/snapshot"
	"github.com/inuoluwadunsimi/backpack/internal/storage"
)

// SystemCollector captures macOS system-level information:
// OS version, Xcode CLI tools, default shell, hardware arch.
// It also populates the SystemTools list with standalone CLI tools
// like go, rustc, java, etc.
type SystemCollector struct{}

func (s *SystemCollector) Name() string { return "system" }

func (s *SystemCollector) IsAvailable() bool {
	return runtime.GOOS == "darwin"
}

func (s *SystemCollector) Collect(manifest *snapshot.ToolsManifest, _ storage.BlobStore) error {
	// TODO: capture sw_vers output (macOS version)
	// TODO: check xcode-select --version
	// TODO: detect standalone tools: go version, rustc --version, java -version
	// TODO: populate manifest.SystemTools
	manifest.SystemTools = []snapshot.SystemToolEntry{}
	return nil
}
