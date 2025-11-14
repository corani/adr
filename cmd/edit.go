package cmd

import (
	"context"
	"log"
	"strconv"

	"github.com/corani/adr/internal/app"
	"github.com/spf13/cobra"
)

// editCmd represents the edit command.
//
//nolint:exhaustruct,gochecknoglobals
var editCmd = &cobra.Command{
	Use:   "edit <id>",
	Short: "open the adr with number <id> in the default editor",
	Run: func(_ *cobra.Command, args []string) {
		ctx := context.TODO()

		number, err := strconv.Atoi(args[0])
		if err != nil {
			log.Printf("invalid argument: %v", err)

			return
		}

		if err := app.Edit(ctx, number); err != nil {
			log.Printf("couldn't edit adr %d: %v", number, err)
		}
	},
}

//nolint:gochecknoinits
func init() {
	rootCmd.AddCommand(editCmd)
}
