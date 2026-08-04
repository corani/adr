package cmd

import (
	"log"
	"os"

	"github.com/corani/adr/internal/app"
	"github.com/spf13/cobra"
)

func NewVersionCommand() *cobra.Command {
	//nolint:exhaustruct
	return &cobra.Command{
		Use:   "version",
		Short: "Show the version information",
		Long:  `Print the build version, commit hash, and build date of the adr binary.`,
		Run: func(_ *cobra.Command, _ []string) {
			if err := app.Version(os.Args[0]); err != nil {
				log.Printf("couldn't show version: %v", err)
			}
		},
	}
}
