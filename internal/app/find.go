package app

import (
	"fmt"
	"regexp"
	"slices"
	"strings"

	"charm.land/glamour/v2"
	"charm.land/lipgloss/v2"
	"github.com/corani/adr/config"
	"github.com/corani/adr/internal/adr"
)

func Find(conf *config.Config, query string, fullText bool) error {
	re := buildQuery(query)

	var rows []string

	err := adr.ForEach(conf, func(entry *adr.Adr) error {
		if matches(re, entry, fullText) {
			rows = append(rows, fmt.Sprintf("| %04d | %s | %s | %s |", entry.Number, entry.Date, entry.Status, entry.Title))
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("%w: find: %w", ErrInternal, err)
	}

	if len(rows) == 0 {
		fmt.Println("no results")

		return nil
	}

	slices.Sort(rows)

	table := "| # | date | status | title |\n|---|------|--------|-------|\n" + strings.Join(rows, "\n") + "\n"

	renderer, err := glamour.NewTermRenderer(
		glamour.WithEnvironmentConfig(),
		glamour.WithWordWrap(0),
	)
	if err != nil {
		return fmt.Errorf("%w: find: %w", ErrInternal, err)
	}

	out, err := renderer.Render(table)
	if err != nil {
		return fmt.Errorf("%w: find: %w", ErrInternal, err)
	}

	if _, err = lipgloss.Print(out); err != nil {
		return fmt.Errorf("%w: find: %w", ErrInternal, err)
	}

	return nil
}

func buildQuery(query string) *regexp.Regexp {
	words := strings.Fields(query)
	parts := make([]string, len(words))

	for i, w := range words {
		parts[i] = regexp.QuoteMeta(w)
	}

	return regexp.MustCompile("(?i)" + strings.Join(parts, ".*"))
}

func matches(pattern *regexp.Regexp, entry *adr.Adr, fullText bool) bool {
	if pattern.MatchString(entry.Title) {
		return true
	}

	if !fullText {
		return false
	}

	return slices.ContainsFunc([]string{
		string(entry.Status),
		entry.Date,
		string(entry.Body),
	}, pattern.MatchString)
}
