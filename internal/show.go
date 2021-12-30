package internal

import (
	"fmt"
	"os"

	markdown "github.com/MichaelMure/go-term-markdown"
	"github.com/corani/adr/internal/adr"
	"github.com/corani/adr/internal/config"
	"github.com/jedib0t/go-pretty/v6/table"
)

func Show(id int) error {
	conf, err := config.ReadConfig()
	if err != nil {
		return err
	}

	found, err := adr.ById(conf, adr.Number(id))
	if err != nil {
		return err
	}

	t := table.NewWriter()

	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleRounded)
	t.AppendRows([]table.Row{
		{"Filename", found.Filename},
		{"Number", fmt.Sprintf("%04d", found.Number)},
		{"Date", found.Date},
		{"Status", found.Status},
	})

	t.Render()

	out := markdown.Render(string(found.Body), 80, 1)

	fmt.Println()
	fmt.Println(string(out))

	return nil
}
