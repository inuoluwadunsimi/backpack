package collectors

import (
	"github.com/inuoluwadunsimi/backpack/internal/snapshot"
	"github.com/inuoluwadunsimi/backpack/internal/storage"
)

// PipCollector captures globally installed pip/pip3 packages and the Python runtime.
type PipCollector struct{}

func (p *PipCollector) Name() string { return "pip" }

func (p *PipCollector) IsAvailable() bool {
	// TODO: check if `pip3` is in PATH
	return false
}

func (p *PipCollector) Collect(manifest *snapshot.ToolsManifest, _ storage.BlobStore) error {
	// TODO: run `python3 --version` to get runtime version
	// TODO: detect manager (pyenv, system)
	// TODO: run `pip3 list --format=json`
	manifest.Python = &snapshot.RuntimeState{}
	manifest.PipPackages = []snapshot.PackageEntry{}
	return nil
}
