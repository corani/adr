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

func InitConfig(path string) (*config.Config, error) {
	root, err := config.ProjectRoot()
	if err != nil {
		return nil, fmt.Errorf("%w: init: %w", ErrInternal, err)
	}

	conf := &config.Config{
		Root:          path,
		AdrTemplate:   filepath.Join(path, "adr-template.md"),
		IndexTemplate: filepath.Join(path, "index-template.md"),
		Project:       root,
	}

	log.Printf(`create ".adr.yaml"`)

	if err := config.WriteConfig(root, conf); err != nil {
		return nil, fmt.Errorf("%w: init: %w", ErrInternal, err)
	}

	return conf, nil
}

func Init(conf *config.Config) error {
	log.Printf("[CMD] mkdir -p %q", conf.Root)

	//nolint:mnd,gofumpt
	if err := os.MkdirAll(filepath.Join(conf.Project, conf.Root), 0o750); err != nil {
		return fmt.Errorf("%w: init: %w", ErrInternal, err)
	}

	log.Printf("create %q", conf.AdrTemplate)

	if err := template.Write("adr.md", filepath.Join(conf.Project, conf.AdrTemplate)); err != nil {
		return fmt.Errorf("%w: init: %w", ErrInternal, err)
	}

	log.Printf("create %q", conf.IndexTemplate)

	if err := template.Write("index.md", filepath.Join(conf.Project, conf.IndexTemplate)); err != nil {
		return fmt.Errorf("%w: init: %w", ErrInternal, err)
	}

	if err := adr.Index(conf); err != nil {
		return fmt.Errorf("%w: init: %w", ErrInternal, err)
	}

	return nil
}
