package internal

import (
	"log"
	"os"
	"path/filepath"

	"github.com/corani/adr/internal/config"
)

func Init(path string) error {
	root, err := config.ProjectRoot()
	if err != nil {
		return err
	}

	conf := &config.Config{
		Root:     path,
		Template: filepath.Join(path, "template.md"),
	}

	log.Printf("[CMD] mkdir -p %q", path)

	if err := os.MkdirAll(filepath.Join(root, path), 0755); err != nil {
		return err
	}

	log.Printf(`create ".adr.yaml"`)

	if err := config.WriteConfig(root, conf); err != nil {
		return err
	}

	log.Printf("create %q", conf.Template)

	if err := writeTemplate(filepath.Join(root, conf.Template)); err != nil {
		return err
	}

	return nil
}
