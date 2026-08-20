package cmd

import (
	"log"

	"github.com/corani/adr/config"
	"github.com/corani/adr/internal/app"
	"github.com/spf13/cobra"
)

func NewInitConfigCommand() *cobra.Command {
	//nolint:exhaustruct_v5
	return &cobra.Command{
		Use:   "init [path]",
		Short: "Initialize the ADR path (default is `docs/adr`)",
		Long: `Initialize the ADR directory and configuration.

Creates the ADR directory, writes .adr.yaml at the project root, and copies
the default ADR and index templates into the directory. The project root is
determined by walking up from the current directory until a .git directory
is found.

The path argument sets the ADR directory relative to the project root
(default: docs/adr).`,
		Args: cobra.MaximumNArgs(1),
		Run: func(_ *cobra.Command, args []string) {
			path := "docs/adr"

			if len(args) > 0 {
				path = args[0]
			}

			conf, err := app.InitConfig(path)
			if err != nil {
				log.Printf("couldn't initialize adr: %v", err)

				return
			}

			if err := app.Init(conf); err != nil {
				log.Printf("couldn't initialize adr: %v", err)
			}
		},
	}
}

func NewInitCommand(conf *config.Config) *cobra.Command {
	//nolint:exhaustruct_v5
	return &cobra.Command{
		Use:   "init",
		Short: "Initialize the ADR path",
		Long: `Initialize the ADR directory using the pre-configured path.

Creates the ADR directory and copies the default ADR and index templates into it.`,
		Run: func(_ *cobra.Command, args []string) {
			if err := app.Init(conf); err != nil {
				log.Printf("couldn't initialize adr: %v", err)
			}
		},
	}
}
