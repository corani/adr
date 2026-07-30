package cmd

import (
	"log"
	"strconv"
	"strings"

	"github.com/corani/adr/config"
	"github.com/corani/adr/internal/app"
	"github.com/spf13/cobra"
)

func NewUpdateCommand(conf *config.Config) *cobra.Command {
	//nolint:exhaustruct
	return &cobra.Command{
		Use:   "update <id> <status>",
		Short: "Update the ADR with number <id> to status <status>",
		Long: `Update the status of the ADR with the given number.

Valid statuses are: proposed, accepted, deprecated, superseded.

The index (README.md) is regenerated automatically after the update.`,
		Args:  cobra.ExactArgs(2), //nolint:mnd
		Run: func(_ *cobra.Command, args []string) {
			number, err := strconv.Atoi(args[0])
			if err != nil {
				log.Printf("invalid argument: %v", err)

				return
			}

			status := strings.ToLower(args[1])

			if err := app.Update(conf, number, status); err != nil {
				log.Printf("couldn't update adr %d: %v", number, err)
			}
		},
	}
}
