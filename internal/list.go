package internal

import (
	"fmt"
	"os"

	"github.com/corani/adr/internal/adr"
	"github.com/jedib0t/go-pretty/v6/table"
)

func List() error {
	root, err := AdrRoot()
	if err != nil {
		return err
	}

	list, err := adr.Index(root)
	if err != nil {
		return err
	}

	t := table.NewWriter()

	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleRounded)
	t.SortBy([]table.SortBy{{Name: "#", Mode: table.AscNumeric}})
	t.AppendHeader(table.Row{"#", "date", "status", "title"})

	for _, v := range list {
		t.AppendRow(table.Row{
			fmt.Sprintf("%04d", v.Number),
			v.Date,
			v.Status,
			v.Title,
		})
	}

	t.Render()

	return nil
}
