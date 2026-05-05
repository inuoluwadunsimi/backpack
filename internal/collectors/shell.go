package collectors

import "github.com/inuoluwadunsimi/backpack/internal/snapshot"

// ShellCollector captures shell configuration files (.zshrc, .bashrc, etc.)
// and the current default shell.
type ShellCollector struct{}

func (s *ShellCollector) Name() string { return "shell" }

func (s *ShellCollector) IsAvailable() bool {
	// Shell is always available on macOS
	return true
}

func (s *ShellCollector) Collect() (*snapshot.ToolState, error) {
	// TODO: detect default shell ($SHELL)
	// TODO: capture dotfiles: .zshrc, .bashrc, .bash_profile, .zprofile,
	//       .gitconfig, .ssh/config (sanitized), etc.
	return &snapshot.ToolState{Name: s.Name()}, nil
}
