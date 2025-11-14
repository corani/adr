package cmd

import (
	"log"
	"os"

	"github.com/corani/adr/internal/app"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command.
//
//nolint:exhaustruct,gochecknoglobals
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show the version information",
	Run: func(_ *cobra.Command, _ []string) {
		if err := app.Version(os.Args[0]); err != nil {
			log.Printf("couldn't show version: %v", err)
		}
	},
}

//nolint:gochecknoinits
func init() {
	rootCmd.AddCommand(versionCmd)
}
