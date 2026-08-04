package app

import (
	"fmt"
	"strings"

	"charm.land/glamour/v2"
	"charm.land/lipgloss/v2"
	"github.com/corani/adr/config"
	"github.com/corani/adr/internal/adr"
)

func Show(conf *config.Config, number int) error {
	found, err := adr.ByID(conf, adr.Number(number))
	if err != nil {
		return fmt.Errorf("%w: show: %w", ErrInternal, err)
	}

	meta := strings.Join([]string{
		"| field | value |",
		"|-------|-------|",
		fmt.Sprintf("| Filename | %s |", found.Filename),
		fmt.Sprintf("| Number | %04d |", found.Number),
		fmt.Sprintf("| Date | %s |", found.Date),
		fmt.Sprintf("| Status | %s |", found.Status),
	}, "\n") + "\n\n" + string(found.Body)

	r, err := glamour.NewTermRenderer(
		glamour.WithEnvironmentConfig(),
		glamour.WithWordWrap(80),
	)
	if err != nil {
		return fmt.Errorf("%w: show: %w", ErrInternal, err)
	}

	out, err := r.Render(meta)
	if err != nil {
		return fmt.Errorf("%w: show: %w", ErrInternal, err)
	}

	lipgloss.Print(out)

	return nil
}
