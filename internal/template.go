package internal

import (
	"embed"
	"fmt"
	"os"
)

//go:embed template
var template embed.FS

func writeTemplate(source, target string) error {
	body, err := template.ReadFile(source)
	if err != nil {
		return fmt.Errorf("%w: writeTemplate: %v", ErrInternal, err)
	}

	//nolint: gosec,gomnd,gofumpt
	if err := os.WriteFile(target, body, 0644); err != nil {
		return fmt.Errorf("%w: writeTemplate: %v", ErrInternal, err)
	}

	return nil
}
