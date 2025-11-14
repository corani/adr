package app

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/corani/adr/internal/adr"
	"github.com/corani/adr/internal/config"
	"github.com/corani/adr/internal/template"
)

func Init(path string) error {
	root, err := config.ProjectRoot()
	if err != nil {
		return fmt.Errorf("%w: init: %w", ErrInternal, err)
	}

	conf := &config.Config{
		Root:          path,
		AdrTemplate:   filepath.Join(path, "adr-template.md"),
		IndexTemplate: filepath.Join(path, "index-template.md"),
		Project:       "",
	}

	log.Printf("[CMD] mkdir -p %q", path)

	//nolint:mnd,gofumpt
	if err := os.MkdirAll(filepath.Join(root, path), 0750); err != nil {
		return fmt.Errorf("%w: init: %w", ErrInternal, err)
	}

	log.Printf(`create ".adr.yaml"`)

	if err := config.WriteConfig(root, conf); err != nil {
		return fmt.Errorf("%w: init: %w", ErrInternal, err)
	}

	log.Printf("create %q", conf.AdrTemplate)

	if err := template.Write("adr.md", filepath.Join(root, conf.AdrTemplate)); err != nil {
		return fmt.Errorf("%w: init: %w", ErrInternal, err)
	}

	log.Printf("create %q", conf.IndexTemplate)

	if err := template.Write("index.md", filepath.Join(root, conf.IndexTemplate)); err != nil {
		return fmt.Errorf("%w: init: %w", ErrInternal, err)
	}

	if err := adr.Index(conf); err != nil {
		return fmt.Errorf("%w: init: %w", ErrInternal, err)
	}

	return nil
}
