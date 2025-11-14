package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands.
//
//nolint:exhaustruct,gochecknoglobals
var rootCmd = &cobra.Command{
	Use:   "adr",
	Short: "A command line tool to maintain Architecture Decision Records",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
