package storage

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
)

// BlobStore provides content-addressable storage for config file contents.
type BlobStore interface {
	// Put stores content and returns its SHA-256 hash.
	// If the hash already exists, this is a no-op (dedup).
	Put(content []byte) (hash string, err error)

	// Get retrieves content by hash.
	Get(hash string) ([]byte, error)

	// Has checks if a blob exists.
	Has(hash string) bool
}

// LocalBlobStore implements BlobStore using the local filesystem.
// Blobs are stored as files named by their SHA-256 hash under a base directory.
type LocalBlobStore struct {
	BasePath string // e.g. ~/.backpack/snapshots/blobs
}

// NewLocalBlobStore creates a LocalBlobStore, ensuring the directory exists.
func NewLocalBlobStore(basePath string) (*LocalBlobStore, error) {
	if err := os.MkdirAll(basePath, 0o755); err != nil {
		return nil, fmt.Errorf("creating blob directory: %w", err)
	}
	return &LocalBlobStore{BasePath: basePath}, nil
}

func (b *LocalBlobStore) Put(content []byte) (string, error) {
	hash := sha256sum(content)
	path := filepath.Join(b.BasePath, hash+".blob")

	// Dedup: skip if blob already exists
	if _, err := os.Stat(path); err == nil {
		return hash, nil
	}

	if err := os.WriteFile(path, content, 0o644); err != nil {
		return "", fmt.Errorf("writing blob %s: %w", hash[:8], err)
	}
	return hash, nil
}

func (b *LocalBlobStore) Get(hash string) ([]byte, error) {
	path := filepath.Join(b.BasePath, hash+".blob")
	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("blob %s not found", hash[:8])
	}
	if err != nil {
		return nil, fmt.Errorf("reading blob %s: %w", hash[:8], err)
	}
	return data, nil
}

func (b *LocalBlobStore) Has(hash string) bool {
	path := filepath.Join(b.BasePath, hash+".blob")
	_, err := os.Stat(path)
	return err == nil
}

// sha256sum returns the hex-encoded SHA-256 hash of data.
func sha256sum(data []byte) string {
	h := sha256.Sum256(data)
	return hex.EncodeToString(h[:])
}
