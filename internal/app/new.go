package app

import (
	"context"
	"fmt"

	"github.com/corani/adr/config"
	"github.com/corani/adr/internal/adr"
)

func Create(ctx context.Context, conf *config.Config, title string) error {
	v, err := adr.Create(conf, title)
	if err != nil {
		return fmt.Errorf("%w: create: %w", ErrInternal, err)
	}

	return Edit(ctx, conf, int(v.Number))
}
