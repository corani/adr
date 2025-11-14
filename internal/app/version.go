package app

import (
	"fmt"

	"github.com/corani/adr/internal/config"
)

func Version(prog string) error {
	version := config.GetVersion()

	fmt.Println(prog)
	fmt.Printf("  version:  %v\n", version.Version)
	fmt.Printf("  commit:   %v\n", version.Commit)
	fmt.Printf("  built at: %v\n", version.BuiltAt)

	return nil
}
