package cmd

import (
	"log"
	"strconv"

	"github.com/corani/adr/config"
	"github.com/corani/adr/internal/app"
	"github.com/spf13/cobra"
)

func NewShowCommand(conf *config.Config) *cobra.Command {
	var format string

	//nolint:exhaustruct_v5
	cmd := &cobra.Command{
		Use:     "show <id>",
		Aliases: []string{"view"},
		Short:   "Show the ADR with number <id>",
		Long: `Show the ADR with the given number.

Prints a summary table (number, date, status, filename) followed by the
rendered markdown body of the ADR.`,
		Args: cobra.ExactArgs(1),
		Run: func(_ *cobra.Command, args []string) {
			number, err := strconv.Atoi(args[0])
			if err != nil {
				log.Printf("invalid id %q: expected a number", args[0])

				return
			}

			outputFormat, err := app.ParseFormat(format)
			if err != nil {
				log.Printf("invalid format %q: must be one of: md, raw, json", format)

				return
			}

			if err := app.Show(conf, number, outputFormat); err != nil {
				log.Printf("couldn't show adr %d: %v", number, err)
			}
		},
	}

	cmd.Flags().StringVarP(&format, "format", "f", string(app.FormatMd), "output format: md, raw, json")

	return cmd
}
