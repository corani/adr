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
		return err
	}

	t := table.NewWriter()

	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleRounded)
	t.SortBy([]table.SortBy{{Name: "#", Mode: table.AscNumeric}})
	t.AppendHeader(table.Row{"#", "date", "status", "title"})

	err = adr.ForEach(conf, func(v *adr.Adr) error {
		t.AppendRow(table.Row{
			fmt.Sprintf("%04d", v.Number),
			v.Date,
			v.Status,
			v.Title,
		})

		return nil
	})
	if err != nil {
		return err
	}

	t.Render()

	return nil
}
