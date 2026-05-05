package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/inuoluwadunsimi/backpack/internal/commands"
)

var version = "0.1.0-dev"

func main() {
	rootCmd := &cobra.Command{
		Use:   "backpack",
		Short: "Back up and restore your entire dev environment",
		Long: `backpack captures all your dev tools, packages, and configs so you can
restore your full development environment on a new machine with a single command.

Supported tools: Homebrew, npm, pip, VS Code extensions, shell configs, and more.`,
		Version: version,
	}

	// Register subcommands
	rootCmd.AddCommand(commands.NewInitCmd())
	rootCmd.AddCommand(commands.NewBackupCmd())
	rootCmd.AddCommand(commands.NewRestoreCmd())
	rootCmd.AddCommand(commands.NewDevicesCmd())
	rootCmd.AddCommand(commands.NewDiffCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
