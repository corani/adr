package cmd

import (
	"log"
	"strings"

	"github.com/corani/adr/internal"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:     "new [title]",
	Aliases: []string{"add", "create"},
	Short:   "create a new ADR with optional title",
	Run: func(cmd *cobra.Command, args []string) {
		title := strings.Join(args, " ")

		if err := internal.Create(title); err != nil {
			log.Printf("couldn't create adr: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
