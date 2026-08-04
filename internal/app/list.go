package app

import (
	"fmt"
	"slices"
	"strings"

	"charm.land/glamour/v2"
	"charm.land/lipgloss/v2"
	"github.com/corani/adr/config"
	"github.com/corani/adr/internal/adr"
)

func List(conf *config.Config) error {
	var rows []string

	err := adr.ForEach(conf, func(v *adr.Adr) error {
		rows = append(rows, fmt.Sprintf("| %04d | %s | %s | %s |", v.Number, v.Date, v.Status, v.Title))
		return nil
	})
	if err != nil {
		return fmt.Errorf("%w: list: %w", ErrInternal, err)
	}

	slices.Sort(rows)

	md := "| # | date | status | title |\n|---|------|--------|-------|\n" + strings.Join(rows, "\n") + "\n"

	r, err := glamour.NewTermRenderer(
		glamour.WithEnvironmentConfig(),
		glamour.WithWordWrap(0),
	)
	if err != nil {
		return fmt.Errorf("%w: list: %w", ErrInternal, err)
	}

	out, err := r.Render(md)
	if err != nil {
		return fmt.Errorf("%w: list: %w", ErrInternal, err)
	}

	lipgloss.Print(out)

	return nil
}
