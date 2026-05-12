package collectors

import (
	"context"

	"github.com/inuoluwadunsimi/backpack/internal/exec"
	"github.com/inuoluwadunsimi/backpack/internal/snapshot"
	"github.com/inuoluwadunsimi/backpack/internal/storage"
)

// ShellCollector captures shell configuration files (.zshrc, .bashrc, etc.),
// the current default shell, parsed aliases, git config, and SSH state.
// Config file contents are stored in the blob store and referenced by hash.
type ShellCollector struct {
	runner exec.Runner
	blobs  storage.BlobStore
}

func NewShellCollector(runner exec.Runner, blobs storage.BlobStore) *ShellCollector {
	return &ShellCollector{runner: runner, blobs: blobs}
}

func (s *ShellCollector) Name() string { return "shell" }

func (s *ShellCollector) Available() bool {
	// Shell is always available on macOS
	return true
}

// ShellCollectorData bundles the multiple manifest fields this collector populates.
type ShellCollectorData struct {
	Shell *snapshot.ShellState
	Git   *snapshot.GitConfigState
	SSH   *snapshot.SSHState
}

func (s *ShellCollector) Collect(ctx context.Context) (*CollectorResult, error) {
	var warnings []string

	// TODO: detect default shell ($SHELL)
	// TODO: read dotfiles: .zshrc, .bashrc, .bash_profile, .zprofile, etc.
	// TODO: for each file, call s.blobs.Put(content) and store the FileRef
	// TODO: parse alias lines from shell config
	// TODO: parse .gitconfig for name, email, editor, aliases
	// TODO: scan ~/.ssh/ for public key files, check for config

	data := &ShellCollectorData{
		Shell: &snapshot.ShellState{
			Type:        "zsh", // TODO: detect from $SHELL
			ConfigFiles: map[string]snapshot.FileRef{},
			Aliases:     []string{},
		},
		Git: &snapshot.GitConfigState{},
		SSH: &snapshot.SSHState{
			KeyFiles: []string{},
		},
	}

	return &CollectorResult{
		Available: true,
		Data:      data,
		Warnings:  warnings,
	}, nil
}
