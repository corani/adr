package app

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"charm.land/glamour/v2"
	"charm.land/lipgloss/v2"
	"github.com/corani/adr/config"
	"github.com/corani/adr/internal/adr"
)

type adrShowEntry struct {
	Number   int    `json:"number"`
	Title    string `json:"title"`
	Status   string `json:"status"`
	Date     string `json:"date"`
	Filepath string `json:"filepath"`
	Body     string `json:"body"`
}

func Show(conf *config.Config, number int, format Format) error {
	found, err := adr.ByID(conf, adr.Number(number))
	if err != nil {
		return fmt.Errorf("%w: show: %w", ErrInternal, err)
	}

	switch format {
	case FormatJSON:
		return showJSON(os.Stdout, found)
	case FormatRaw, FormatMd:
		return showMarkdown(os.Stdout, found, format)
	default:
		return fmt.Errorf("%w: show: unknown format %q: must be one of: md, raw, json", ErrInternal, format)
	}
}

func showJSON(writer io.Writer, found *adr.Adr) error {
	entry := adrShowEntry{
		Number:   int(found.Number),
		Title:    found.Title,
		Status:   string(found.Status),
		Date:     found.Date,
		Filepath: found.Filename,
		Body:     strings.TrimSpace(string(found.Body)),
	}

	out, err := json.MarshalIndent(entry, "", "  ")
	if err != nil {
		return fmt.Errorf("%w: show: %w", ErrInternal, err)
	}

	if _, err = fmt.Fprintln(writer, string(out)); err != nil {
		return fmt.Errorf("%w: show: %w", ErrInternal, err)
	}

	return nil
}

func showMarkdown(writer io.Writer, found *adr.Adr, format Format) error {
	meta := strings.Join([]string{
		"| field | value |",
		"|-------|-------|",
		fmt.Sprintf("| Filename | %s |", found.Filename),
		fmt.Sprintf("| Number | %04d |", found.Number),
		fmt.Sprintf("| Date | %s |", found.Date),
		fmt.Sprintf("| Status | %s |", found.Status),
	}, "\n") + "\n\n" + reflowBody(string(found.Body))

	if format == FormatRaw {
		if _, err := fmt.Fprint(writer, meta); err != nil {
			return fmt.Errorf("%w: show: %w", ErrInternal, err)
		}

		return nil
	}

	renderer, err := glamour.NewTermRenderer(
		glamour.WithEnvironmentConfig(),
		glamour.WithWordWrap(0),
	)
	if err != nil {
		return fmt.Errorf("%w: show: %w", ErrInternal, err)
	}

	out, err := renderer.Render(meta)
	if err != nil {
		return fmt.Errorf("%w: show: %w", ErrInternal, err)
	}

	if _, err = lipgloss.Print(out); err != nil {
		return fmt.Errorf("%w: show: %w", ErrInternal, err)
	}

	return nil
}
