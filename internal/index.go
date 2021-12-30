package internal

import (
	"github.com/corani/adr/internal/adr"
	"github.com/corani/adr/internal/config"
)

func Index() error {
	conf, err := config.ReadConfig()
	if err != nil {
		return err
	}

	body, err := template.ReadFile("template/index.md")
	if err != nil {
		return err
	}

	return adr.Index(conf, string(body))
}
