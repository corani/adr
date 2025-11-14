package cmd

import (
	"log"

	"github.com/corani/adr/internal/app"
	"github.com/spf13/cobra"
)

// listCmd represents the list command.
//
//nolint:exhaustruct,gochecknoglobals
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all ADRs with their id, date and status",
	Run: func(_ *cobra.Command, _ []string) {
		if err := app.List(); err != nil {
			log.Printf("couldn't list adrs: %v", err)
		}
	},
}

//nolint:gochecknoinits
func init() {
	rootCmd.AddCommand(listCmd)
}
