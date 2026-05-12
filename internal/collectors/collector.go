package collectors

import (
	"github.com/inuoluwadunsimi/backpack/internal/exec"
	"github.com/inuoluwadunsimi/backpack/internal/storage"
)

// Registry returns all known collectors in dependency order.
// The Runner and BlobStore are injected so collectors can execute
// commands and store config file content.
func Registry(runner exec.Runner, blobs storage.BlobStore) []Collector {
	return []Collector{
		NewSystemCollector(runner),
		NewHomebrewCollector(runner),
		NewShellCollector(runner, blobs),
		NewNPMCollector(runner),
		NewPipCollector(runner),
		NewVSCodeCollector(runner),
	}
}
