package internal

import (
	"fmt"
	"os"

	markdown "github.com/MichaelMure/go-term-markdown"
	"github.com/corani/adr/internal/adr"
	"github.com/corani/adr/internal/config"
	"github.com/jedib0t/go-pretty/v6/table"
)

func Show(number int) error {
	conf, err := config.ReadConfig()
	if err != nil {
		return fmt.Errorf("%w: show: %v", ErrInternal, err)
	}

	found, err := adr.ByID(conf, adr.Number(number))
	if err != nil {
		return fmt.Errorf("%w: show: %v", ErrInternal, err)
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
