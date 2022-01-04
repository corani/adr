package cmd

import (
	"log"
	"strconv"
	"strings"

	"github.com/corani/adr/internal"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command.
//nolint:exhaustivestruct,gochecknoglobals
var updateCmd = &cobra.Command{
	Use:   "update <id> <status>",
	Short: "update the adr with number <id> to status <status>",
	Args:  cobra.ExactArgs(2), //nolint:gomnd
	Run: func(cmd *cobra.Command, args []string) {
		number, err := strconv.Atoi(args[0])
		if err != nil {
			log.Printf("invalid argument: %v", err)

			return
		}

		status := strings.ToLower(args[1])

		if err := internal.Update(number, status); err != nil {
			log.Printf("couldn't update adr %d: %v", number, err)
		}
	},
}

//nolint:gochecknoinits
func init() {
	rootCmd.AddCommand(updateCmd)
}
