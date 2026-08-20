package cmd

import (
	"context"
	"log"
	"strings"

	"github.com/corani/adr/config"
	"github.com/corani/adr/internal/app"
	"github.com/spf13/cobra"
)

func NewNewCommand(conf *config.Config) *cobra.Command {
	//nolint:exhaustruct_v5
	return &cobra.Command{
		Use:     "new [title]",
		Aliases: []string{"add", "create"},
		Short:   "Create a new ADR with optional title",
		Long: `Create a new ADR with an auto-incremented number and optional title.

The title can be supplied as arguments or left empty to be filled in later.
After creating the file, it is immediately opened in the default editor
(see $EDITOR).`,
		Run: func(_ *cobra.Command, args []string) {
			ctx := context.TODO()
			title := strings.Join(args, " ")

			if err := app.Create(ctx, conf, title); err != nil {
				log.Printf("couldn't create adr: %v", err)
			}
		},
	}
}
