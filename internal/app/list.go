package app

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"slices"

	"github.com/corani/adr/config"
	"github.com/corani/adr/internal/adr"
)

func List(conf *config.Config, format Format) error {
	var entries []adrListEntry

	err := adr.ForEach(conf, func(entry *adr.Adr) error {
		entries = append(entries, adrListEntry{
			Number:   int(entry.Number),
			Title:    entry.Title,
			Status:   string(entry.Status),
			Date:     entry.Date,
			Filepath: entry.Filename,
		})

		return nil
	})
	if err != nil {
		return fmt.Errorf("%w: list: %w", ErrInternal, err)
	}

	slices.SortFunc(entries, func(a, b adrListEntry) int {
		return a.Number - b.Number
	})

	switch format {
	case FormatJSON:
		return listJSON(os.Stdout, entries)
	case FormatRaw, FormatMd:
		return renderMarkdownTable(os.Stdout, entries, format, "list")
	default:
		return fmt.Errorf("%w: list: unknown format %q: must be one of: md, raw, json", ErrInternal, format)
	}
}

func listJSON(writer io.Writer, entries []adrListEntry) error {
	out, err := json.MarshalIndent(entries, "", "  ")
	if err != nil {
		return fmt.Errorf("%w: list: %w", ErrInternal, err)
	}

	if _, err = fmt.Fprintln(writer, string(out)); err != nil {
		return fmt.Errorf("%w: list: %w", ErrInternal, err)
	}

	return nil
}
