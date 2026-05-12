package storage

import "github.com/inuoluwadunsimi/backpack/internal/snapshot"

// S3Backend stores snapshots in an AWS S3 bucket.
type S3Backend struct {
	Bucket    string
	Region    string
	Prefix    string
	Profile   string
	blobStore *LocalBlobStore
}

func (s *S3Backend) Blobs() BlobStore {
	return s.blobStore
}

func (s *S3Backend) SaveSnapshot(snap *snapshot.Snapshot) error {
	// TODO: marshal to JSON and upload to s3://<bucket>/<prefix>/<deviceID>/<snapshotID>.json
	return nil
}

func (s *S3Backend) LoadSnapshot(id string) (*snapshot.Snapshot, error) {
	return nil, nil
}

func (s *S3Backend) ListSnapshots(deviceID string) ([]snapshot.Snapshot, error) {
	return nil, nil
}

func (s *S3Backend) DeleteSnapshot(id string) error {
	return nil
}

func (s *S3Backend) SaveDevice(device *snapshot.Device) error {
	return nil
}

func (s *S3Backend) ListDevices() ([]snapshot.Device, error) {
	return nil, nil
}

func (s *S3Backend) GetDevice(id string) (*snapshot.Device, error) {
	return nil, nil
}
