// Package storage defines the Backend interface for persisting snapshots
// and provides implementations for local filesystem, git, and S3.
package storage

import "github.com/inuoluwadunsimi/backpack/internal/snapshot"

// Backend abstracts snapshot persistence.
type Backend interface {
	// SaveSnapshot persists a snapshot. Overwrites if same ID exists.
	SaveSnapshot(snap *snapshot.Snapshot) error

	// LoadSnapshot retrieves a snapshot by ID.
	LoadSnapshot(id string) (*snapshot.Snapshot, error)

	// ListSnapshots returns all snapshot IDs for a given device, newest first.
	ListSnapshots(deviceID string) ([]snapshot.Snapshot, error)

	// DeleteSnapshot removes a snapshot by ID.
	DeleteSnapshot(id string) error

	// SaveDevice persists device metadata.
	SaveDevice(device *snapshot.Device) error

	// ListDevices returns all registered devices.
	ListDevices() ([]snapshot.Device, error)

	// GetDevice retrieves a device by ID.
	GetDevice(id string) (*snapshot.Device, error)

	// Blobs returns the blob store for config file content.
	Blobs() BlobStore
}
