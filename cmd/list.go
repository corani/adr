package cmd

import (
	"log"

	"github.com/corani/adr/internal"
	"github.com/spf13/cobra"
)

// listCmd represents the list command.
//nolint:exhaustivestruct,gochecknoglobals
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all ADRs with their id, date and status",
	Run: func(cmd *cobra.Command, args []string) {
		if err := internal.List(); err != nil {
			log.Printf("couldn't list adrs: %v", err)
		}
	},
}

//nolint:gochecknoinits
func init() {
	rootCmd.AddCommand(listCmd)
}
