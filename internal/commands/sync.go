package commands

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/inuoluwadunsimi/backpack/internal/config"
)

// NewSyncCmd creates the `backpack sync` command.
func NewSyncCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sync",
		Short: "Push latest local snapshot to remote storage",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := config.Load()
			if err != nil {
				return fmt.Errorf("loading config: %w", err)
			}

			if cfg.Device.ID == "" {
				return fmt.Errorf("backpack not initialized — run `backpack init` first")
			}

			switch cfg.Storage.Backend {
			case config.BackendGit:
				// TODO: implement sync to git
				fmt.Println("Syncing to git remote...")

			case config.BackendS3:
				// TODO: implement sync to s3
				fmt.Println("Syncing to S3...")

			case config.BackendLocal:
				fmt.Println("Storage backend is local — nothing to sync.")

			default:
				return fmt.Errorf("unknown storage backend: %s", cfg.Storage.Backend)
			}

			return nil
		},
	}

	return cmd
}