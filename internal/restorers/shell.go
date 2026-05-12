package restorers

import "github.com/inuoluwadunsimi/backpack/internal/snapshot"

type ShellRestorer struct{}

func (s *ShellRestorer) Name() string { return "shell" }

func (s *ShellRestorer) Restore(manifest *snapshot.ToolsManifest, dryRun bool) error {
	if manifest.Shell == nil {
		return nil
	}
	// TODO: restore dotfiles (.zshrc, .bashrc, .gitconfig, etc.)
	// TODO: retrieve config file content from blob store via FileRef.Hash
	// TODO: prompt before overwriting existing files
	return nil
}
