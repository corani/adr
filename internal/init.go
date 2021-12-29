package internal

import (
	"log"
	"os"
	"path/filepath"
)

func Init(path string) error {
	root, err := ProjectRoot()
	if err != nil {
		return err
	}

	path = filepath.Join(root, path)

	log.Printf("[CMD] mkdir -p %q", path)

	if err := os.MkdirAll(path, 0755); err != nil {
		return err
	}

	log.Println("create `template.md`")

	if err := initTemplate(filepath.Join(path, "template.md")); err != nil {
		return err
	}

	return nil
}

func initTemplate(path string) error {
	body := []byte(`---
type: adr
number: {{.Number}}
title: {{.Title}}
date: {{.Date}}
status: {{.Status}}
link: {{.Link}}
---
# {{.Number}}. {{.Title}}

# Context

# Decision

# Consequences

`)

	return os.WriteFile(path, body, 0644)
}
