package internal

import (
	"fmt"

	markdown "github.com/MichaelMure/go-term-markdown"
	"github.com/corani/adr/internal/adr"
)

func Show(id int) error {
	root, err := AdrRoot()
	if err != nil {
		return err
	}

	list, err := adr.Index(root)
	if err != nil {
		return err
	}

	for _, v := range list {
		if v.Number == adr.Number(id) {
			out := markdown.Render(string(v.Body), 80, 2)

			fmt.Println(string(out))

			return nil
		}
	}

	return fmt.Errorf("file not found")
}
