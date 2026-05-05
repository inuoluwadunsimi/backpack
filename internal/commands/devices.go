package commands

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/inuoluwadunsimi/backpack/internal/config"
	"github.com/inuoluwadunsimi/backpack/internal/storage"
)

// NewDevicesCmd creates the `backpack devices` command.
func NewDevicesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "devices",
		Short: "List all registered devices",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := config.Load()
			if err != nil {
				return fmt.Errorf("loading config: %w", err)
			}

			backend, err := storage.NewLocalBackend(cfg.Storage.LocalPath)
			if err != nil {
				return fmt.Errorf("initializing storage: %w", err)
			}

			devices, err := backend.ListDevices()
			if err != nil {
				return fmt.Errorf("listing devices: %w", err)
			}

			if len(devices) == 0 {
				fmt.Println("No devices registered yet. Run `backpack init` to get started.")
				return nil
			}

			fmt.Printf("Registered devices (%d):\n\n", len(devices))
			for _, d := range devices {
				current := ""
				if d.ID == cfg.DeviceID {
					current = " ← this machine"
				}
				fmt.Printf("  %s  %s%s\n", d.ID[:8], d.Name, current)
				fmt.Printf("         OS: %s/%s  |  Hostname: %s\n", d.OS, d.Arch, d.Hostname)
				fmt.Printf("         Created: %s  |  Last seen: %s\n\n",
					d.Created.Format("2006-01-02"), d.LastSeen.Format("2006-01-02 15:04"))
			}

			return nil
		},
	}

	return cmd
}
