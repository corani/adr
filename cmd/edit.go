package cmd

import (
	"context"
	"log"
	"strconv"

	"github.com/corani/adr/config"
	"github.com/corani/adr/internal/app"
	"github.com/spf13/cobra"
)

func NewEditCommand(conf *config.Config) *cobra.Command {
	//nolint:exhaustruct_v5
	return &cobra.Command{
		Use:   "edit <id>",
		Short: "Open the ADR with number <id> in the default editor",
		Long: `Open the ADR with the given number in the default editor.

The editor is determined by the $EDITOR environment variable. If $EDITOR is
not set, the command will fail.`,
		Args: cobra.ExactArgs(1),
		Run: func(_ *cobra.Command, args []string) {
			ctx := context.TODO()

			number, err := strconv.Atoi(args[0])
			if err != nil {
				log.Printf("invalid id %q: expected a number", args[0])

				return
			}

			if err := app.Edit(ctx, conf, number); err != nil {
				log.Printf("couldn't edit adr %d: %v", number, err)
			}
		},
	}
}
