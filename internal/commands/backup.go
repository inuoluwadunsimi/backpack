package commands

import (
	"fmt"
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

			if cfg.DeviceID == "" {
				return fmt.Errorf("backpack not initialized — run `backpack init` first")
			}

			backend, err := storage.NewLocalBackend(cfg.Storage.LocalPath)
			if err != nil {
				return fmt.Errorf("initializing storage: %w", err)
			}

			// Collect state from all available tools
			snap := &snapshot.Snapshot{
				ID:        uuid.New().String(),
				DeviceID:  cfg.DeviceID,
				Timestamp: time.Now(),
				Tools:     make(map[string]snapshot.ToolState),
				Meta: snapshot.SnapshotMeta{
					Trigger: "manual",
					Notes:   note,
					Version: "0.1.0", // TODO: inject at build time
				},
			}

			for _, c := range collectors.Registry() {
				if !c.IsAvailable() {
					fmt.Printf("  ⏭  %s — not installed, skipping\n", c.Name())
					continue
				}

				fmt.Printf("  📦 Collecting %s...\n", c.Name())
				state, err := c.Collect()
				if err != nil {
					fmt.Printf("  ⚠  %s — error: %v\n", c.Name(), err)
					continue
				}
				snap.Tools[c.Name()] = *state
			}

			if err := backend.SaveSnapshot(snap); err != nil {
				return fmt.Errorf("saving snapshot: %w", err)
			}

			fmt.Printf("\n✓ Snapshot saved: %s\n", snap.ID[:8])
			fmt.Printf("  Tools captured: %d\n", len(snap.Tools))

			return nil
		},
	}

	cmd.Flags().StringVarP(&note, "note", "m", "", "optional note for this snapshot")

	return cmd
}
