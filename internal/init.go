package internal

import (
	"log"
	"os"
	"path/filepath"

	"github.com/corani/adr/internal/adr"
	"github.com/corani/adr/internal/config"
)

func Init(path string) error {
	root, err := config.ProjectRoot()
	if err != nil {
		return err
	}

	conf := &config.Config{
		Root:          path,
		AdrTemplate:   filepath.Join(path, "adr-template.md"),
		IndexTemplate: filepath.Join(path, "index-template.md"),
		Project:       "",
	}

	log.Printf("[CMD] mkdir -p %q", path)

	//nolint:gomnd,gofumpt
	if err := os.MkdirAll(filepath.Join(root, path), 0755); err != nil {
		return err
	}

	log.Printf(`create ".adr.yaml"`)

	if err := config.WriteConfig(root, conf); err != nil {
		return err
	}

	log.Printf("create %q", conf.AdrTemplate)

	if err := writeTemplate(
		"template/template.md", filepath.Join(root, conf.AdrTemplate),
	); err != nil {
		return err
	}

	log.Printf("create %q", conf.IndexTemplate)

	if err := writeTemplate(
		"template/index.md", filepath.Join(root, conf.IndexTemplate),
	); err != nil {
		return err
	}

	return adr.Index(conf)
}
