package commands

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/inuoluwadunsimi/backpack/internal/config"
	"github.com/inuoluwadunsimi/backpack/internal/storage"
)

// NewRestoreCmd creates the `backpack restore` command.
func NewRestoreCmd() *cobra.Command {
	var (
		deviceID   string
		snapshotID string
		dryRun     bool
		toolFilter []string
	)

	cmd := &cobra.Command{
		Use:   "restore",
		Short: "Restore dev tools from a snapshot",
		Long:  "Restores packages and configs from a previous snapshot. You can select which device and snapshot to restore from.",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := config.Load()
			if err != nil {
				return fmt.Errorf("loading config: %w", err)
			}

			backend, err := storage.NewLocalBackend(cfg.LocalPath())
			if err != nil {
				return fmt.Errorf("initializing storage: %w", err)
			}

			// If no device specified, list available devices
			if deviceID == "" {
				devices, err := backend.ListDevices()
				if err != nil {
					return fmt.Errorf("listing devices: %w", err)
				}
				if len(devices) == 0 {
					return fmt.Errorf("no devices found — run `backpack init` on a machine first")
				}

				fmt.Println("Available devices:")
				for _, d := range devices {
					fmt.Printf("  • %s (%s) — %s/%s\n", d.Name, d.ID[:8], d.OS, d.Arch)
				}
				fmt.Println("\nUse --device <id> to select a device")
				return nil
			}

			// If no snapshot specified, show latest for that device
			if snapshotID == "" {
				snapshots, err := backend.ListSnapshots(deviceID)
				if err != nil {
					return fmt.Errorf("listing snapshots: %w", err)
				}
				if len(snapshots) == 0 {
					return fmt.Errorf("no snapshots found for device %s", deviceID)
				}

				fmt.Println("Available snapshots:")
				for _, s := range snapshots {
					fmt.Printf("  • %s — %s\n",
						s.ID[:8], s.CapturedAt.Format("2006-01-02 15:04"))
				}
				fmt.Println("\nUse --snapshot <id> to select a snapshot")
				return nil
			}

			// TODO: load the snapshot and run restorers in dependency order
			_ = dryRun
			_ = toolFilter
			fmt.Printf("Restoring from snapshot %s (dry-run: %v)...\n", snapshotID, dryRun)

			return nil
		},
	}

	cmd.Flags().StringVarP(&deviceID, "device", "d", "", "device ID to restore from")
	cmd.Flags().StringVarP(&snapshotID, "snapshot", "s", "", "snapshot ID to restore (defaults to latest)")
	cmd.Flags().BoolVar(&dryRun, "dry-run", false, "show what would be restored without making changes")
	cmd.Flags().StringSliceVarP(&toolFilter, "tools", "t", nil, "only restore specific tools (e.g. --tools homebrew,npm)")

	return cmd
}
