package template

import (
	"embed"
	"fmt"
)

//go:embed template.md index.md
var files embed.FS

func Get(name string) ([]byte, error) {
	b, err := files.ReadFile(name)
	if err != nil {
		return nil, fmt.Errorf("read embedded template %s: %w", name, err)
	}

	return b, nil
}
