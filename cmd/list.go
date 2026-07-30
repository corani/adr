package cmd

import (
	"log"

	"github.com/corani/adr/internal/app"
	"github.com/spf13/cobra"
)

func NewListCommand() *cobra.Command {
	//nolint:exhaustruct
	return &cobra.Command{
		Use:   "list",
		Short: "list all ADRs with their id, date and status",
		Run: func(_ *cobra.Command, _ []string) {
			if err := app.List(); err != nil {
				log.Printf("couldn't list adrs: %v", err)
			}
		},
	}
}
