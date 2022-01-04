package cmd

import (
	"log"
	"os"

	"github.com/corani/adr/internal"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show the version information",
	Run: func(cmd *cobra.Command, args []string) {
		if err := internal.Version(os.Args[0]); err != nil {
			log.Printf("couldn't show version: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
