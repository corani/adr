package app

import (
	"fmt"
	"io"
	"strings"

	"charm.land/glamour/v2"
	"charm.land/lipgloss/v2"
)

func renderMarkdownTable(writer io.Writer, entries []adrListEntry, format Format, label string) error {
	rows := make([]string, len(entries))

	for i, e := range entries {
		rows[i] = fmt.Sprintf("| %04d | %s | %s | %s |", e.Number, e.Date, e.Status, e.Title)
	}

	table := "| # | date | status | title |\n|---|------|--------|-------|\n" + strings.Join(rows, "\n") + "\n\n" +
		"_Use `adr show <id>` to view a specific ADR._\n"

	if format == FormatRaw {
		if _, err := fmt.Fprint(writer, table); err != nil {
			return fmt.Errorf("%w: %s: %w", ErrInternal, label, err)
		}

		return nil
	}

	renderer, err := glamour.NewTermRenderer(
		glamour.WithEnvironmentConfig(),
		glamour.WithWordWrap(0),
	)
	if err != nil {
		return fmt.Errorf("%w: %s: %w", ErrInternal, label, err)
	}

	rendered, err := renderer.Render(table)
	if err != nil {
		return fmt.Errorf("%w: %s: %w", ErrInternal, label, err)
	}

	if _, err = lipgloss.Print(rendered); err != nil {
		return fmt.Errorf("%w: %s: %w", ErrInternal, label, err)
	}

	return nil
}
