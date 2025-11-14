package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/corani/adr/internal/adr"
	"github.com/corani/adr/internal/config"
)

func Edit(ctx context.Context, number int) error {
	conf, err := config.ReadConfig()
	if err != nil {
		return fmt.Errorf("%w: edit: %w", ErrInternal, err)
	}

	found, err := adr.ByID(conf, adr.Number(number))
	if err != nil {
		return fmt.Errorf("%w: edit: %w", ErrInternal, err)
	}

	log.Printf("editing ADR: %v", filepath.Join(conf.Root, found.Filename))

	// #nosec G204
	cmd := exec.CommandContext(ctx, os.Getenv("EDITOR"), filepath.Join(conf.Project, conf.Root, found.Filename))
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
