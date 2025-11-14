package cmd

import (
	"context"
	"log"
	"strings"

	"github.com/corani/adr/internal/app"
	"github.com/spf13/cobra"
)

// newCmd represents the new command.
//
//nolint:exhaustruct,gochecknoglobals
var newCmd = &cobra.Command{
	Use:     "new [title]",
	Aliases: []string{"add", "create"},
	Short:   "create a new ADR with optional title",
	Run: func(_ *cobra.Command, args []string) {
		ctx := context.TODO()
		title := strings.Join(args, " ")

		if err := app.Create(ctx, title); err != nil {
			log.Printf("couldn't create adr: %v", err)
		}
	},
}

//nolint:gochecknoinits
func init() {
	rootCmd.AddCommand(newCmd)
}
