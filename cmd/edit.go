package cmd

import (
	"log"
	"strconv"

	"github.com/corani/adr/internal"
	"github.com/spf13/cobra"
)

// editCmd represents the edit command.
//nolint:exhaustivestruct,gochecknoglobals
var editCmd = &cobra.Command{
	Use:   "edit <id>",
	Short: "open the adr with number <id> in the default editor",
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			log.Printf("invalid argument: %v", err)

			return
		}

		if err := internal.Edit(id); err != nil {
			log.Printf("couldn't edit adr %d: %v", id, err)
		}
	},
}

//nolint:gochecknoinits
func init() {
	rootCmd.AddCommand(editCmd)
}
