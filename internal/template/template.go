package template

import (
	"fmt"
	"os"
)

func Write(source, target string) error {
	body, err := Get(source)
	if err != nil {
		return fmt.Errorf("%w: writeTemplate: %w", ErrInternal, err)
	}

	//nolint:mnd,gofumpt
	if err := os.WriteFile(target, body, 0600); err != nil {
		return fmt.Errorf("%w: writeTemplate: %w", ErrInternal, err)
	}

	return nil
}
