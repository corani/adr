package app

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
	"slices"
	"strings"

	"github.com/corani/adr/config"
	"github.com/corani/adr/internal/adr"
)

func Find(conf *config.Config, query string, fullText bool, format Format) error {
	re := buildQuery(query)

	var entries []adrListEntry

	err := adr.ForEach(conf, func(entry *adr.Adr) error {
		if matches(re, entry, fullText) {
			entries = append(entries, adrListEntry{
				Number:   int(entry.Number),
				Title:    entry.Title,
				Status:   string(entry.Status),
				Date:     entry.Date,
				Filepath: entry.Filename,
			})
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("%w: find: %w", ErrInternal, err)
	}

	if len(entries) == 0 {
		fmt.Println("no results")

		return nil
	}

	slices.SortFunc(entries, func(a, b adrListEntry) int {
		return a.Number - b.Number
	})

	switch format {
	case FormatJSON:
		return findJSON(os.Stdout, entries)
	case FormatRaw, FormatMd:
		return renderMarkdownTable(os.Stdout, entries, format, "find")
	default:
		return fmt.Errorf("%w: find: unknown format %q: must be one of: md, raw, json", ErrInternal, format)
	}
}

func findJSON(writer io.Writer, entries []adrListEntry) error {
	out, err := json.MarshalIndent(entries, "", "  ")
	if err != nil {
		return fmt.Errorf("%w: find: %w", ErrInternal, err)
	}

	if _, err = fmt.Fprintln(writer, string(out)); err != nil {
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
