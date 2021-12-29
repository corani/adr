package internal

import (
	"github.com/corani/adr/internal/adr"
)

func Create(title string) error {
	root, err := AdrRoot()
	if err != nil {
		return err
	}

	if v, err := adr.Create(root, title); err != nil {
		return err
	} else {
		Edit(int(v.Number))
	}

	return nil
}
