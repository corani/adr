package cmd

import (
	"log"

	"github.com/corani/adr/internal"
	"github.com/spf13/cobra"
)

// initCmd represents the init command.
//
//nolint:exhaustruct,gochecknoglobals
var initCmd = &cobra.Command{
	Use:   "init [path]",
	Short: "initialize the adr path (default is `docs/adr`)",
	Args:  cobra.MaximumNArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		path := "docs/adr"

		if len(args) > 0 {
			path = args[0]
		}

		if err := internal.Init(path); err != nil {
			log.Printf("couldn't initialize adr: %v", err)

			return
		}
	},
}

//nolint:gochecknoinits
func init() {
	rootCmd.AddCommand(initCmd)
}
