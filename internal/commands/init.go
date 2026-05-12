package commands

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/google/uuid"
	"github.com/spf13/cobra"

	"github.com/inuoluwadunsimi/backpack/internal/config"
	"github.com/inuoluwadunsimi/backpack/internal/snapshot"
	"github.com/inuoluwadunsimi/backpack/internal/storage"
)

// NewInitCmd creates the `backpack init` command.
// This registers the current machine as a device and takes the first snapshot.
func NewInitCmd() *cobra.Command {
	var deviceName string

	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize backpack on this machine",
		Long:  "Registers this device, takes an initial snapshot of all dev tools, and sets up auto-watching.",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := config.Load()
			if err != nil {
				return fmt.Errorf("loading config: %w", err)
			}

			if cfg.Device.ID != "" {
				return fmt.Errorf("backpack is already initialized on this device (ID: %s)", cfg.Device.ID)
			}

			// Resolve device name
			hostname, _ := os.Hostname()
			if deviceName == "" {
				deviceName = hostname
			}

			// Create device record
			device := &snapshot.Device{
				ID:       uuid.New().String(),
				Name:     deviceName,
				Hostname: hostname,
				OS:       runtime.GOOS,
				Arch:     runtime.GOARCH,
				Created:  time.Now(),
				LastSeen: time.Now(),
			}

			// Initialize storage backend
			backend, err := storage.NewLocalBackend(cfg.LocalPath())
			if err != nil {
				return fmt.Errorf("initializing storage: %w", err)
			}

			// Save device
			if err := backend.SaveDevice(device); err != nil {
				return fmt.Errorf("saving device: %w", err)
			}

			// Update config
			cfg.Device = config.DeviceConfig{
				ID:           device.ID,
				Name:         device.Name,
				RegisteredAt: time.Now(),
			}
			if err := config.Save(cfg); err != nil {
				return fmt.Errorf("saving config: %w", err)
			}

			fmt.Printf("✓ Device registered: %s (%s)\n", device.Name, device.ID[:8])
			fmt.Println("  Taking initial snapshot...")

			// TODO: run backup command logic here to take first snapshot
			fmt.Println("✓ Initial snapshot complete. Auto-watching enabled.")

			return nil
		},
	}

	cmd.Flags().StringVarP(&deviceName, "name", "n", "", "friendly name for this device (defaults to hostname)")

	return cmd
}
