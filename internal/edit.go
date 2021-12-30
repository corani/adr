package internal

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/corani/adr/internal/adr"
	"github.com/corani/adr/internal/config"
)

func Edit(id int) error {
	conf, err := config.ReadConfig()
	if err != nil {
		return err
	}

	list, err := adr.Index(conf)
	if err != nil {
		return err
	}

	for _, v := range list {
		if v.Number == adr.Number(id) {
			cmd := exec.Command(os.Getenv("EDITOR"), filepath.Join(conf.Project, conf.Root, v.Filename))
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			return cmd.Run()
		}
	}

	return fmt.Errorf("file not found")
}
