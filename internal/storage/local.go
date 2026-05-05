package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/inuoluwadunsimi/backpack/internal/snapshot"
)

// LocalBackend stores snapshots as JSON files on the local filesystem.
type LocalBackend struct {
	BasePath string // e.g. ~/.backpack/snapshots
}

func NewLocalBackend(basePath string) (*LocalBackend, error) {
	if err := os.MkdirAll(basePath, 0o755); err != nil {
		return nil, fmt.Errorf("creating storage directory: %w", err)
	}
	return &LocalBackend{BasePath: basePath}, nil
}

func (l *LocalBackend) snapshotDir(deviceID string) string {
	return filepath.Join(l.BasePath, "devices", deviceID, "snapshots")
}

func (l *LocalBackend) deviceDir() string {
	return filepath.Join(l.BasePath, "devices")
}

func (l *LocalBackend) SaveSnapshot(snap *snapshot.Snapshot) error {
	dir := l.snapshotDir(snap.DeviceID)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(snap, "", "  ")
	if err != nil {
		return err
	}

	path := filepath.Join(dir, snap.ID+".json")
	return os.WriteFile(path, data, 0o644)
}

func (l *LocalBackend) LoadSnapshot(id string) (*snapshot.Snapshot, error) {
	// TODO: search across all device directories for the snapshot ID
	return nil, fmt.Errorf("LoadSnapshot not yet implemented")
}

func (l *LocalBackend) ListSnapshots(deviceID string) ([]snapshot.Snapshot, error) {
	dir := l.snapshotDir(deviceID)
	entries, err := os.ReadDir(dir)
	if os.IsNotExist(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	var snapshots []snapshot.Snapshot
	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".json" {
			continue
		}
		data, err := os.ReadFile(filepath.Join(dir, entry.Name()))
		if err != nil {
			continue
		}
		var snap snapshot.Snapshot
		if err := json.Unmarshal(data, &snap); err != nil {
			continue
		}
		snapshots = append(snapshots, snap)
	}

	// Sort newest first
	sort.Slice(snapshots, func(i, j int) bool {
		return snapshots[i].Timestamp.After(snapshots[j].Timestamp)
	})

	return snapshots, nil
}

func (l *LocalBackend) DeleteSnapshot(id string) error {
	// TODO: find and delete the snapshot file
	return fmt.Errorf("DeleteSnapshot not yet implemented")
}

func (l *LocalBackend) SaveDevice(device *snapshot.Device) error {
	dir := filepath.Join(l.deviceDir(), device.ID)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(device, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filepath.Join(dir, "device.json"), data, 0o644)
}

func (l *LocalBackend) ListDevices() ([]snapshot.Device, error) {
	dir := l.deviceDir()
	entries, err := os.ReadDir(dir)
	if os.IsNotExist(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	var devices []snapshot.Device
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		data, err := os.ReadFile(filepath.Join(dir, entry.Name(), "device.json"))
		if err != nil {
			continue
		}
		var dev snapshot.Device
		if err := json.Unmarshal(data, &dev); err != nil {
			continue
		}
		devices = append(devices, dev)
	}

	return devices, nil
}

func (l *LocalBackend) GetDevice(id string) (*snapshot.Device, error) {
	path := filepath.Join(l.deviceDir(), id, "device.json")
	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("device %s not found", id)
	}
	if err != nil {
		return nil, err
	}

	var dev snapshot.Device
	if err := json.Unmarshal(data, &dev); err != nil {
		return nil, err
	}
	return &dev, nil
}
