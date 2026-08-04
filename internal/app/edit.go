package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"golang.org/x/term"

	"github.com/corani/adr/config"
	"github.com/corani/adr/internal/adr"
)

func Edit(ctx context.Context, conf *config.Config, number int) error {
	found, err := adr.ByID(conf, adr.Number(number))
	if err != nil {
		return fmt.Errorf("%w: edit: %w", ErrInternal, err)
	}

	path := filepath.Join(conf.Project, conf.Root, found.Filename)

	if !term.IsTerminal(int(os.Stdout.Fd())) {
		log.Printf("ADR saved at: %s", path)

		return nil
	}

	editor := os.Getenv("VISUAL")
	if editor == "" {
		editor = os.Getenv("EDITOR")
	}

	log.Printf("editing ADR: %v", filepath.Join(conf.Root, found.Filename))

	// #nosec G204,G702 // Command injection via environment variable
	cmd := exec.CommandContext(ctx, editor, path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("%w: edit: %w", ErrInternal, err)
	}

	if err := adr.Index(conf); err != nil {
		return fmt.Errorf("%w: edit: %w", ErrInternal, err)
	}

	return nil
}
