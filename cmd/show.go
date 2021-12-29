package cmd

import (
	"log"
	"strconv"

	"github.com/corani/adr/internal"
	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show <id>",
	Short: "show the adr with number <id>",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			log.Printf("invalid argument: %v", err)

			return
		}

		if err := internal.Show(id); err != nil {
			log.Printf("couldn't show adr %d: %v", id, err)
		}
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
