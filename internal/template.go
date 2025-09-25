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
		return fmt.Errorf("%w: writeTemplate: %w", ErrInternal, err)
	}

	//nolint:mnd,gofumpt
	if err := os.WriteFile(target, body, 0600); err != nil {
		return fmt.Errorf("%w: writeTemplate: %w", ErrInternal, err)
	}

	return nil
}
