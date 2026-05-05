// Package watcher monitors for new package installations and auto-triggers
// snapshots after the initial setup. It watches for changes to key files
// like Brewfile.lock, global npm list, etc.
package watcher

import (
	"fmt"
	"os"
	"path/filepath"
)

// WatchPaths returns the filesystem paths that should be monitored for
// changes to detect new package installations.
func WatchPaths() []string {
	home, _ := os.UserHomeDir()
	return []string{
		// Homebrew
		"/usr/local/Cellar",
		"/opt/homebrew/Cellar",
		// npm global
		filepath.Join(home, ".npm"),
		// pip
		filepath.Join(home, "Library/Python"),
		// VS Code extensions
		filepath.Join(home, ".vscode/extensions"),
		// Shell configs
		filepath.Join(home, ".zshrc"),
		filepath.Join(home, ".bashrc"),
		filepath.Join(home, ".gitconfig"),
	}
}

// Watcher monitors filesystem changes and triggers snapshots.
type Watcher struct {
	Paths    []string
	OnChange func() error // callback to trigger a snapshot
}

// NewWatcher creates a new filesystem watcher.
func NewWatcher(onChange func() error) *Watcher {
	return &Watcher{
		Paths:    WatchPaths(),
		OnChange: onChange,
	}
}

// Start begins watching. This blocks until Stop is called or an error occurs.
func (w *Watcher) Start() error {
	// TODO: use fsnotify to watch paths
	// TODO: debounce changes (wait 5s after last change before triggering)
	// TODO: call w.OnChange() when a meaningful change is detected
	return fmt.Errorf("watcher not yet implemented")
}

// Stop terminates the watcher.
func (w *Watcher) Stop() error {
	// TODO: close fsnotify watcher
	return nil
}
