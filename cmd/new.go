package cmd

import (
	"log"
	"strings"

	"github.com/corani/adr/internal"
	"github.com/spf13/cobra"
)

// newCmd represents the new command.
//nolint:exhaustivestruct,gochecknoglobals
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

//nolint:gochecknoinits
func init() {
	rootCmd.AddCommand(newCmd)
}
