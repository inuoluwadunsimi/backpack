package commands

import (
	"context"
	"fmt"
	"runtime"
	"time"

	"github.com/google/uuid"
	"github.com/spf13/cobra"

	"github.com/inuoluwadunsimi/backpack/internal/collectors"
	"github.com/inuoluwadunsimi/backpack/internal/config"
	bpexec "github.com/inuoluwadunsimi/backpack/internal/exec"
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

			runner := bpexec.NewRunner()
			ctx := context.Background()

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
			for _, c := range collectors.Registry(runner, backend.Blobs()) {
				result, err := collectors.SafeCollect(ctx, c, collectors.DefaultTimeout)
				if err != nil {
					fmt.Printf("  ⚠  %s — error: %v\n", c.Name(), err)
					continue
				}

				if !result.Available {
					fmt.Printf("  ⏭  %s — not installed, skipping\n", c.Name())
					continue
				}

				fmt.Printf("  📦 Collected %s (%s)\n", c.Name(), result.Duration.Round(time.Millisecond))
				for _, w := range result.Warnings {
					fmt.Printf("     ⚠ %s\n", w)
				}

				applyCollectorResult(&snap.Tools, c.Name(), result)
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

// applyCollectorResult assigns typed collector data to the right manifest fields.
func applyCollectorResult(m *snapshot.ToolsManifest, name string, result *collectors.CollectorResult) {
	switch data := result.Data.(type) {
	case *snapshot.HomebrewState:
		m.Homebrew = data
	case *snapshot.VSCodeState:
		m.VSCode = data
	case []snapshot.SystemToolEntry:
		m.SystemTools = data
	case *collectors.ShellCollectorData:
		m.Shell = data.Shell
		m.Git = data.Git
		m.SSH = data.SSH
	case *collectors.NPMCollectorData:
		m.Node = data.Node
		m.NpmGlobals = data.NpmGlobals
	case *collectors.PipCollectorData:
		m.Python = data.Python
		m.PipPackages = data.PipPackages
	}
}
