package internal

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/corani/adr/internal/adr"
	"github.com/corani/adr/internal/config"
)

func Edit(number int) error {
	conf, err := config.ReadConfig()
	if err != nil {
		return err
	}

	found, err := adr.ByID(conf, adr.Number(number))
	if err != nil {
		return err
	}

	log.Printf("editing ADR: %v", filepath.Join(conf.Root, found.Filename))

	//nolint: gosec
	cmd := exec.Command(os.Getenv("EDITOR"), filepath.Join(conf.Project, conf.Root, found.Filename))
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	return adr.Index(conf)
}
