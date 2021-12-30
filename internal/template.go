package internal

import (
	"embed"
	"fmt"
	"os"
)

//go:embed template
var template embed.FS

func writeTemplate(source, target string) error {
	body, err := template.ReadFile(source)
	if err != nil {
		return fmt.Errorf("ReadFile: %w", err)
	}

	return os.WriteFile(target, body, 0644)
}
