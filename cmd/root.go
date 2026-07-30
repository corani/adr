package cmd

import (
	"github.com/corani/adr/config"
	"github.com/spf13/cobra"
)

// EmbedCommands returns the commands suitable for embedding in another CLI.
// The caller is responsible for providing a fully populated conf; init-config
// and version are excluded as they are specific to the standalone adr CLI.
func EmbedCommands(conf *config.Config) []*cobra.Command {
	return []*cobra.Command{
		NewInitCommand(conf),
		NewNewCommand(conf),
		NewListCommand(conf),
		NewShowCommand(conf),
		NewEditCommand(conf),
		NewUpdateCommand(conf),
	}
}

// AdrCommands returns the full command set for the standalone adr CLI,
// including init-config (which discovers the project root and writes .adr.yaml)
// and version.
func AdrCommands(conf *config.Config) []*cobra.Command {
	return []*cobra.Command{
		NewInitConfigCommand(),
		NewVersionCommand(),
		NewNewCommand(conf),
		NewListCommand(conf),
		NewShowCommand(conf),
		NewEditCommand(conf),
		NewUpdateCommand(conf),
	}
}
