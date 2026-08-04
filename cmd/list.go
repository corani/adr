package cmd

import (
	"log"

	"github.com/corani/adr/config"
	"github.com/corani/adr/internal/app"
	"github.com/spf13/cobra"
)

func NewListCommand(conf *config.Config) *cobra.Command {
	var format string

	//nolint:exhaustruct
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all ADRs with their id, date and status",
		Run: func(_ *cobra.Command, _ []string) {
			outputFormat, err := app.ParseFormat(format)
			if err != nil {
				log.Printf("invalid format %q: %v", format, err)

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
