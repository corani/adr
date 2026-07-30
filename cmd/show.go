package cmd

import (
	"log"
	"strconv"

	"github.com/corani/adr/config"
	"github.com/corani/adr/internal/app"
	"github.com/spf13/cobra"
)

func NewShowCommand(conf *config.Config) *cobra.Command {
	//nolint:exhaustruct
	return &cobra.Command{
		Use:     "show <id>",
		Aliases: []string{"view"},
		Short:   "Show the ADR with number <id>",
		Long: `Show the ADR with the given number.

Prints a summary table (number, date, status, filename) followed by the
rendered markdown body of the ADR.`,
		Args:    cobra.ExactArgs(1),
		Run: func(_ *cobra.Command, args []string) {
			number, err := strconv.Atoi(args[0])
			if err != nil {
				log.Printf("invalid argument: %v", err)

				return
			}

			if err := app.Show(conf, number); err != nil {
				log.Printf("couldn't show adr %d: %v", number, err)
			}
		},
	}
}
