package collectors

import (
	"github.com/inuoluwadunsimi/backpack/internal/snapshot"
	"github.com/inuoluwadunsimi/backpack/internal/storage"
)

// NPMCollector captures globally installed npm packages and the Node.js runtime.
type NPMCollector struct{}

func (n *NPMCollector) Name() string { return "npm" }

func (n *NPMCollector) IsAvailable() bool {
	// TODO: check if `npm` is in PATH
	return false
}

func (n *NPMCollector) Collect(manifest *snapshot.ToolsManifest, _ storage.BlobStore) error {
	// TODO: run `node --version` to get runtime version
	// TODO: detect manager (nvm, fnm, system)
	// TODO: run `npm list -g --json --depth=0` for global packages
	manifest.Node = &snapshot.RuntimeState{}
	manifest.NpmGlobals = []snapshot.PackageEntry{}
	return nil
}
