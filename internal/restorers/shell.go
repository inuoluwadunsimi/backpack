package restorers

import "github.com/inuoluwadunsimi/backpack/internal/snapshot"

type ShellRestorer struct{}

func (s *ShellRestorer) Name() string { return "shell" }

func (s *ShellRestorer) Restore(state *snapshot.ToolState, dryRun bool) error {
	// TODO: restore dotfiles (.zshrc, .bashrc, .gitconfig, etc.)
	// TODO: prompt before overwriting existing files
	return nil
}
