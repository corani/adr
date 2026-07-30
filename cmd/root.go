package cmd

import (
	"github.com/spf13/cobra"
)

func NewRootCommand(name string) *cobra.Command {
	//nolint:exhaustruct
	root := &cobra.Command{
		Use:   name,
		Short: "A command line tool to maintain Architecture Decision Records",
	}

	root.AddCommand(
		NewInitCommand(),
		NewNewCommand(),
		NewListCommand(),
		NewShowCommand(),
		NewEditCommand(),
		NewUpdateCommand(),
		NewVersionCommand(),
	)

	return root
}

// ADRCommands returns the commands suitable for embedding in another CLI.
func ADRCommands() []*cobra.Command {
	return []*cobra.Command{
		NewInitCommand(),
		NewNewCommand(),
		NewListCommand(),
		NewShowCommand(),
		NewEditCommand(),
		NewUpdateCommand(),
	}
}
