package app

import (
	"errors"
	"fmt"

	"github.com/corani/adr/config"
	"github.com/corani/adr/internal/adr"
)

var ErrInvalidStatus = errors.New("invalid status")

func Update(conf *config.Config, number int, status string) error {
	switch adr.Status(status) {
	case adr.StatusProposed, adr.StatusAccepted,
		adr.StatusDeprecated, adr.StatusSuperseded:
		// okay
	default:
		return fmt.Errorf("%w: %q: must be one of: proposed, accepted, deprecated, superseded", ErrInvalidStatus, status)
	}

	found, err := adr.ByID(conf, adr.Number(number))
	if err != nil {
		return fmt.Errorf("%w: update: %w", ErrInternal, err)
	}

	found.Status = adr.Status(status)

	if err := adr.Update(conf, found); err != nil {
		return fmt.Errorf("%w: update: %w", ErrInternal, err)
	}

	if err := adr.Index(conf); err != nil {
		return fmt.Errorf("%w: update: %w", ErrInternal, err)
	}

	return nil
}
