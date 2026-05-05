package storage

import "github.com/inuoluwadunsimi/backpack/internal/snapshot"

// GitBackend stores snapshots in a git repository, enabling version history
// and remote sync (push/pull).
type GitBackend struct {
	RepoPath string
	Branch   string
}

func (g *GitBackend) SaveSnapshot(snap *snapshot.Snapshot) error {
	// TODO: write snapshot JSON to repo, git add, git commit
	return nil
}

func (g *GitBackend) LoadSnapshot(id string) (*snapshot.Snapshot, error) {
	// TODO: read from repo
	return nil, nil
}

func (g *GitBackend) ListSnapshots(deviceID string) ([]snapshot.Snapshot, error) {
	// TODO: list from repo
	return nil, nil
}

func (g *GitBackend) DeleteSnapshot(id string) error {
	return nil
}

func (g *GitBackend) SaveDevice(device *snapshot.Device) error {
	return nil
}

func (g *GitBackend) ListDevices() ([]snapshot.Device, error) {
	return nil, nil
}

func (g *GitBackend) GetDevice(id string) (*snapshot.Device, error) {
	return nil, nil
}
