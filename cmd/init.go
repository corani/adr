package cmd

import (
	"log"

	"github.com/corani/adr/config"
	"github.com/corani/adr/internal/app"
	"github.com/spf13/cobra"
)

func NewInitConfigCommand() *cobra.Command {
	//nolint:exhaustruct
	return &cobra.Command{
		Use:   "init [path]",
		Short: "initialize the adr path (default is `docs/adr`)",
		Args:  cobra.MaximumNArgs(1),
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
	//nolint:exhaustruct
	return &cobra.Command{
		Use:   "init",
		Short: "initialize the adr path",
		Run: func(_ *cobra.Command, args []string) {
			if err := app.Init(conf); err != nil {
				log.Printf("couldn't initialize adr: %v", err)
			}
		},
	}
}
