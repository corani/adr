package internal

import (
	"fmt"

	"github.com/corani/adr/internal/adr"
	"github.com/corani/adr/internal/config"
)

func Update(id int, status string) error {
	switch adr.Status(status) {
	case adr.StatusProposed, adr.StatusAccepted, adr.StatusDeprecated, adr.StatusSuperseded:
		// ok
	default:
		return fmt.Errorf("invalid status: %v", status)
	}

	conf, err := config.ReadConfig()
	if err != nil {
		return err
	}

	found, err := adr.ById(conf, adr.Number(id))
	if err != nil {
		return err
	}

	found.Status = adr.Status(status)

	if err := adr.Update(conf, found); err != nil {
		return err
	}

	return adr.Index(conf)
}
