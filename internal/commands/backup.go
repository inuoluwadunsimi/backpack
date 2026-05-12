package commands

import (
	"fmt"
	"runtime"
	"time"

	"github.com/google/uuid"
	"github.com/spf13/cobra"

	"github.com/inuoluwadunsimi/backpack/internal/collectors"
	"github.com/inuoluwadunsimi/backpack/internal/config"
	"github.com/inuoluwadunsimi/backpack/internal/snapshot"
	"github.com/inuoluwadunsimi/backpack/internal/storage"
)

// NewBackupCmd creates the `backpack backup` command.
func NewBackupCmd() *cobra.Command {
	var note string

	cmd := &cobra.Command{
		Use:   "backup",
		Short: "Take a snapshot of all dev tools on this machine",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := config.Load()
			if err != nil {
				return fmt.Errorf("loading config: %w", err)
			}

			if cfg.Device.ID == "" {
				return fmt.Errorf("backpack not initialized — run `backpack init` first")
			}

			backend, err := storage.NewLocalBackend(cfg.LocalPath())
			if err != nil {
				return fmt.Errorf("initializing storage: %w", err)
			}

			// Build snapshot shell
			snap := &snapshot.Snapshot{
				SchemaVersion: 1,
				ID:            uuid.New().String(),
				CapturedAt:    time.Now(),
				Trigger:       "manual",
				Notes:         note,
				Device: snapshot.SnapshotDevice{
					ID:   cfg.Device.ID,
					Name: cfg.Device.Name,
					OS:   runtime.GOOS,
					Arch: runtime.GOARCH,
					// TODO: populate OSVersion from sw_vers
				},
				Tools: snapshot.ToolsManifest{},
			}

			// Collect state from all available collectors
			for _, c := range collectors.Registry() {
				if !c.IsAvailable() {
					fmt.Printf("  ⏭  %s — not installed, skipping\n", c.Name())
					continue
				}

				fmt.Printf("  📦 Collecting %s...\n", c.Name())
				if err := c.Collect(&snap.Tools, backend.Blobs()); err != nil {
					fmt.Printf("  ⚠  %s — error: %v\n", c.Name(), err)
					continue
				}
			}

			if err := backend.SaveSnapshot(snap); err != nil {
				return fmt.Errorf("saving snapshot: %w", err)
			}

			fmt.Printf("\n✓ Snapshot saved: %s\n", snap.ID[:8])

			return nil
		},
	}

	cmd.Flags().StringVarP(&note, "note", "m", "", "optional note for this snapshot")

	return cmd
}
