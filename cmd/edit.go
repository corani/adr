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
	//nolint:exhaustruct
	return &cobra.Command{
		Use:   "edit <id>",
		Short: "open the adr with number <id> in the default editor",
		Args:  cobra.ExactArgs(1),
		Run: func(_ *cobra.Command, args []string) {
			ctx := context.TODO()

			number, err := strconv.Atoi(args[0])
			if err != nil {
				log.Printf("invalid argument: %v", err)

				return
			}

			if err := app.Edit(ctx, conf, number); err != nil {
				log.Printf("couldn't edit adr %d: %v", number, err)
			}
		},
	}
}
