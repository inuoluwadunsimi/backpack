package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewDiffCmd creates the `backpack diff` command.
func NewDiffCmd() *cobra.Command {
	var (
		snapshotA string
		snapshotB string
	)

	cmd := &cobra.Command{
		Use:   "diff",
		Short: "Show differences between snapshots or current state",
		Long:  "Compare two snapshots, or compare a snapshot against the current machine state.",
		RunE: func(cmd *cobra.Command, args []string) error {
			// TODO: load snapshots or collect current state
			// TODO: compute diffs per tool
			// TODO: render diff output

			if snapshotA == "" {
				fmt.Println("Comparing current state against latest snapshot...")
			} else if snapshotB == "" {
				fmt.Printf("Comparing snapshot %s against current state...\n", snapshotA)
			} else {
				fmt.Printf("Comparing snapshot %s vs %s...\n", snapshotA, snapshotB)
			}

			fmt.Println("\n(diff logic not yet implemented)")
			return nil
		},
	}

	cmd.Flags().StringVarP(&snapshotA, "from", "a", "", "first snapshot ID (defaults to latest)")
	cmd.Flags().StringVarP(&snapshotB, "to", "b", "", "second snapshot ID (defaults to current state)")

	return cmd
}
