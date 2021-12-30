package internal

import (
	"github.com/corani/adr/internal/adr"
	"github.com/corani/adr/internal/config"
)

func Create(title string) error {
	conf, err := config.ReadConfig()
	if err != nil {
		return err
	}

	if v, err := adr.Create(conf, title); err != nil {
		return err
	} else {
		Edit(int(v.Number))
	}

	return nil
}
