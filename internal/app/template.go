package app

import (
	"fmt"
	"os"

	tmpl "github.com/corani/adr/internal/template"
)

func writeTemplate(source, target string) error {
	body, err := tmpl.Get(source)
	if err != nil {
		return fmt.Errorf("%w: writeTemplate: %w", ErrInternal, err)
	}

	//nolint:mnd,gofumpt
	if err := os.WriteFile(target, body, 0600); err != nil {
		return fmt.Errorf("%w: writeTemplate: %w", ErrInternal, err)
	}

	return nil
}
