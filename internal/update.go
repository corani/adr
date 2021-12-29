package internal

import (
	"fmt"

	"github.com/corani/adr/internal/adr"
)

func Update(id int, status string) error {
	switch adr.Status(status) {
	case adr.StatusProposed, adr.StatusAccepted, adr.StatusDeprecated, adr.StatusSuperseded:
		// ok
	default:
		return fmt.Errorf("invalid status: %v", status)
	}

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
			v.Status = adr.Status(status)

			return adr.Update(root, v)
		}
	}

	return fmt.Errorf("file not found")
}
