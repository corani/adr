package cmd

import (
	"log"
	"strconv"

	"github.com/corani/adr/internal"
	"github.com/spf13/cobra"
)

// showCmd represents the show command.
//
//nolint:exhaustruct,gochecknoglobals
var showCmd = &cobra.Command{
	Use:     "show <id>",
	Aliases: []string{"view"},
	Short:   "show the adr with number <id>",
	Args:    cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		number, err := strconv.Atoi(args[0])
		if err != nil {
			log.Printf("invalid argument: %v", err)

			return
		}

		if err := internal.Show(number); err != nil {
			log.Printf("couldn't show adr %d: %v", number, err)
		}
	},
}

//nolint:gochecknoinits
func init() {
	rootCmd.AddCommand(showCmd)
}
