package cmd

import (
	"log"

	"github.com/corani/adr/config"
	"github.com/corani/adr/internal/app"
	"github.com/spf13/cobra"
)

func NewListCommand(conf *config.Config) *cobra.Command {
	var format string

	//nolint:exhaustruct_v5
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all ADRs with their id, date and status",
		Long: `List all ADRs in the configured directory.

Prints a table with each ADR's number, date, status, and title, sorted by
number. Use --format to control output: md (rendered, default), raw (plain
markdown table), or json (machine-readable).`,
		Run: func(_ *cobra.Command, _ []string) {
			outputFormat, err := app.ParseFormat(format)
			if err != nil {
				log.Printf("invalid format %q: must be one of: md, raw, json", format)

				return
			}

			if err := app.List(conf, outputFormat); err != nil {
				log.Printf("couldn't list adrs: %v", err)
			}
		},
	}

	cmd.Flags().StringVarP(&format, "format", "f", string(app.FormatMd), "output format: md, raw, json")

	return cmd
}
