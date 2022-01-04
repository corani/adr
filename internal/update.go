package internal

import (
	"errors"
	"fmt"

	"github.com/corani/adr/internal/adr"
	"github.com/corani/adr/internal/config"
)

var ErrInvalidStatus = errors.New("invalid status")

func Update(number int, status string) error {
	switch adr.Status(status) {
	case adr.StatusProposed, adr.StatusAccepted, adr.StatusDeprecated, adr.StatusSuperseded:
		// ok
	default:
		return fmt.Errorf("%w: %v", ErrInvalidStatus, status)
	}

	conf, err := config.ReadConfig()
	if err != nil {
		return err
	}

	found, err := adr.ByID(conf, adr.Number(number))
	if err != nil {
		return err
	}

	found.Status = adr.Status(status)

	if err := adr.Update(conf, found); err != nil {
		return err
	}

	return adr.Index(conf)
}
