package collectors

import (
	"github.com/inuoluwadunsimi/backpack/internal/snapshot"
	"github.com/inuoluwadunsimi/backpack/internal/storage"
)

// ShellCollector captures shell configuration files (.zshrc, .bashrc, etc.),
// the current default shell, and parsed aliases.
// Config file contents are stored in the blob store and referenced by hash.
type ShellCollector struct{}

func (s *ShellCollector) Name() string { return "shell" }

func (s *ShellCollector) IsAvailable() bool {
	// Shell is always available on macOS
	return true
}

func (s *ShellCollector) Collect(manifest *snapshot.ToolsManifest, blobs storage.BlobStore) error {
	// TODO: detect default shell ($SHELL)
	// TODO: read dotfiles: .zshrc, .bashrc, .bash_profile, .zprofile,
	//       .gitconfig, .ssh/config (sanitized), etc.
	// TODO: for each file, call blobs.Put(content) and store the FileRef
	// TODO: parse alias lines from shell config
	manifest.Shell = &snapshot.ShellState{
		Type:        "zsh", // TODO: detect from $SHELL
		ConfigFiles: map[string]snapshot.FileRef{},
		Aliases:     []string{},
	}

	// TODO: populate manifest.Git from .gitconfig
	manifest.Git = &snapshot.GitConfigState{}

	// TODO: populate manifest.SSH from ~/.ssh/
	manifest.SSH = &snapshot.SSHState{
		KeyFiles: []string{},
	}

	return nil
}
