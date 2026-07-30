package app

import (
	"fmt"
	"os"

	markdown "github.com/MichaelMure/go-term-markdown"
	"github.com/corani/adr/config"
	"github.com/corani/adr/internal/adr"
	"github.com/jedib0t/go-pretty/v6/table"
)

func Show(conf *config.Config, number int) error {
	found, err := adr.ByID(conf, adr.Number(number))
	if err != nil {
		return fmt.Errorf("%w: show: %w", ErrInternal, err)
	}

	tbl := table.NewWriter()

	tbl.SetOutputMirror(os.Stdout)
	tbl.SetStyle(table.StyleRounded)
	tbl.AppendRows([]table.Row{
		{"Filename", found.Filename},
		{"Number", fmt.Sprintf("%04d", found.Number)},
		{"Date", found.Date},
		{"Status", found.Status},
	})

	tbl.Render()

	const (
		width  = 80
		indent = 1
	)

	out := markdown.Render(string(found.Body), width, indent)

	fmt.Println()
	fmt.Println(string(out))

	return nil
}
