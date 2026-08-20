package cmd

import (
	"log"

	"github.com/corani/adr/config"
	"github.com/corani/adr/internal/app"
	"github.com/spf13/cobra"
)

func NewFindCommand(conf *config.Config) *cobra.Command {
	var fullText bool

	var format string

	//nolint:exhaustruct_v5
	cmd := &cobra.Command{
		Use:   "find <query>",
		Short: "Find ADRs matching a query",
		Long: `Find ADRs whose title matches the query.

Words in the query are matched in order with anything allowed between them,
so "my search term" matches any title containing "my", then "search", then "term".

Use --text to also search frontmatter fields and the body.`,
		Args: cobra.ExactArgs(1),
		Run: func(_ *cobra.Command, args []string) {
			outputFormat, err := app.ParseFormat(format)
			if err != nil {
				log.Printf("invalid format %q: must be one of: md, raw, json", format)

				return
			}

			if err := app.Find(conf, args[0], fullText, outputFormat); err != nil {
				log.Printf("couldn't find adrs: %v", err)
			}
		},
	}

	cmd.Flags().BoolVarP(&fullText, "text", "t", false, "also search frontmatter fields and body")
	cmd.Flags().StringVarP(&format, "format", "f", string(app.FormatMd), "output format: md, raw, json")

	return cmd
}
