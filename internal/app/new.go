package app

import (
	"context"
	"fmt"

	"github.com/corani/adr/internal/adr"
	"github.com/corani/adr/internal/config"
)

func Create(ctx context.Context, title string) error {
	conf, err := config.ReadConfig()
	if err != nil {
		return fmt.Errorf("%w: create: %w", ErrInternal, err)
	}

	v, err := adr.Create(conf, title)
	if err != nil {
		return fmt.Errorf("%w: create: %w", ErrInternal, err)
	}

	return Edit(ctx, int(v.Number))
}
