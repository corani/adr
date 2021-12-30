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

	list, err := adr.Index(conf)
	if err != nil {
		return err
	}

	for _, v := range list {
		if v.Number == adr.Number(id) {
			t := table.NewWriter()

			t.SetOutputMirror(os.Stdout)
			t.SetStyle(table.StyleRounded)
			t.AppendRows([]table.Row{
				{"Filename", v.Filename},
				{"Number", fmt.Sprintf("%04d", v.Number)},
				{"Date", v.Date},
				{"Status", v.Status},
			})

			t.Render()

			out := markdown.Render(string(v.Body), 80, 1)

			fmt.Println()
			fmt.Println(string(out))

			return nil
		}
	}

	return fmt.Errorf("file not found")
}
