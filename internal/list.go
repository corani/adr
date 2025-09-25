package internal

import (
	"fmt"
	"os"

	"github.com/corani/adr/internal/adr"
	"github.com/corani/adr/internal/config"
	"github.com/jedib0t/go-pretty/v6/table"
)

func List() error {
	conf, err := config.ReadConfig()
	if err != nil {
		return fmt.Errorf("%w: list: %w", ErrInternal, err)
	}

	tbl := table.NewWriter()

	tbl.SetOutputMirror(os.Stdout)
	tbl.SetStyle(table.StyleRounded)
	tbl.SortBy([]table.SortBy{{
		Name:   "#",
		Number: 0,
		Mode:   table.AscNumeric,
	}})
	tbl.AppendHeader(table.Row{"#", "date", "status", "title"})

	err = adr.ForEach(conf, func(v *adr.Adr) error {
		tbl.AppendRow(table.Row{
			fmt.Sprintf("%04d", v.Number),
			v.Date,
			v.Status,
			v.Title,
		})

		return nil
	})
	if err != nil {
		return fmt.Errorf("%w: list: %w", ErrInternal, err)
	}

	tbl.Render()

	return nil
}
